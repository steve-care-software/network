package sqllites

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	token_dashboards "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/dashboards"
	commands_layers "steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/schemas"
	"steve.care/network/domain/schemas/roots"
	"steve.care/network/domain/schemas/roots/groups"
	schema_group_methods "steve.care/network/domain/schemas/roots/groups/methods"
	schema_resources "steve.care/network/domain/schemas/roots/groups/resources"
	field_types "steve.care/network/domain/schemas/roots/groups/resources/fields/types"
	schema_root_methods "steve.care/network/domain/schemas/roots/methods"
)

type resourceRepository struct {
	hashAdapter      hash.Adapter
	signatureAdapter signers.SignatureAdapter
	builder          resources.Builder
	tokenBuilder     tokens.Builder
	dashboardBuilder token_dashboards.Builder
	viewportBuilder  viewports.Builder
	cmdLayerBuilder  commands_layers.LayerBuilder
	builders         map[string]interface{}
	schema           schemas.Schema
	dbPtr            *sql.DB
}

func createResourceRepository(
	hashAdapter hash.Adapter,
	signatureAdapter signers.SignatureAdapter,
	builder resources.Builder,
	tokenBuilder tokens.Builder,
	dashboardBuilder token_dashboards.Builder,
	viewportBuilder viewports.Builder,
	cmdLayerBuilder commands_layers.LayerBuilder,
	builders map[string]interface{},
	schema schemas.Schema,
	dbPtr *sql.DB,
) resources.Repository {
	out := resourceRepository{
		hashAdapter:      hashAdapter,
		signatureAdapter: signatureAdapter,
		builder:          builder,
		tokenBuilder:     tokenBuilder,
		dashboardBuilder: dashboardBuilder,
		viewportBuilder:  viewportBuilder,
		cmdLayerBuilder:  cmdLayerBuilder,
		builders:         builders,
		schema:           schema,
		dbPtr:            dbPtr,
	}

	return &out
}

// Amount returns the amount of resources
func (app *resourceRepository) Amount() (uint, error) {
	return 0, nil
}

// AmountByQuery returns the amount of resources by criteria
func (app *resourceRepository) AmountByQuery(criteria hash.Hash) (uint, error) {
	return 0, nil
}

// ListByQuery lists resource hashes by criteria
func (app *resourceRepository) ListByQuery(criteria hash.Hash) ([]hash.Hash, error) {
	return nil, nil
}

// RetrieveByQuery retrieves a resource by criteria
func (app *resourceRepository) RetrieveByQuery(criteria hash.Hash) (resources.Resource, error) {
	return nil, nil
}

// RetrieveByHash retrieves a resource by hash
func (app *resourceRepository) RetrieveByHash(hash hash.Hash) (resources.Resource, error) {
	rows, err := app.dbPtr.Query("SELECT token, signature FROM resource WHERE hash = ?", hash.Bytes())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		str := fmt.Sprintf("the given hash (%s) do NOT match a Layer instance", hash.String())
		return nil, errors.New(str)
	}

	var retSignatureBytes []byte
	var retTokenHashBytes []byte
	err = rows.Scan(&retTokenHashBytes, &retSignatureBytes)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	pTokenHash, err := app.hashAdapter.FromBytes(retTokenHashBytes)
	if err != nil {
		return nil, err
	}

	token, err := app.retrieveTokenByHash(*pTokenHash, app.schema.Root())
	if err != nil {
		return nil, err
	}

	signature, err := app.signatureAdapter.ToSignature(retSignatureBytes)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithToken(token).
		WithSignature(signature).
		Now()
}

func (app *resourceRepository) retrieveTokenByHash(
	hash hash.Hash,
	root roots.Root,
) (tokens.Token, error) {
	fieldNames, values, err := app.fetchRetrievalFields(root)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT %s, created_on FROM token WHERE hash = ?", strings.Join(fieldNames, ","))
	rows, err := app.dbPtr.Query(query, hash.Bytes())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		str := fmt.Sprintf("the given hash (%s) do NOT match a Token instance", hash.String())
		return nil, errors.New(str)
	}

	var createdOnString string
	values = append(values, &createdOnString)
	err = rows.Scan(values...)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	createdOn, err := time.Parse(timeLayout, createdOnString)
	if err != nil {
		return nil, err
	}

	for idx, oneFieldName := range fieldNames {
		value := values[idx]
		if value == nil {
			continue
		}

		keynames := []string{}
		parentName := ""
		var resourceSchema schema_resources.Resource
		var currentGroup groups.Group
		nextElements := []groups.Element{}
		groupMethods := []schema_group_methods.Methods{}
		rootMethods := []schema_root_methods.Methods{}
		names := strings.Split(oneFieldName, groupNameDelimiterForTableNames)
		for _, oneName := range names {
			if currentGroup == nil {
				if root.Name() == oneName {

					// reuse this
					nextElements = []groups.Element{}
					rootMethods = append(rootMethods, root.Methods())
					chainList := root.Chains().List()
					for _, oneChain := range chainList {
						nextElements = append(nextElements, oneChain.Element())
					}

					parentName = oneName
					keynames = append(keynames, parentName)
					continue
				}
			}

			for _, oneElement := range nextElements {
				if oneElement.IsGroup() {
					group := oneElement.Group()
					if group.Name() == oneName {
						currentGroup = group

						// reuse this
						nextElements = []groups.Element{}
						currentGroupMethods := currentGroup.Methods()
						groupMethods = append(groupMethods, currentGroupMethods)
						rootMethods = append(rootMethods, currentGroupMethods)
						chainList := currentGroup.Chains().List()
						for _, oneChain := range chainList {
							nextElements = append(nextElements, oneChain.Element())
						}

						parentName = fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, oneName)
						keynames = append(keynames, parentName)
						continue
					}

					// error
				}

				if oneElement.IsResource() {
					resource := oneElement.Resource()
					if resource.Name() == oneName {
						resourceSchema = resource
						groupMethods = append(groupMethods, resourceSchema.Builder())

						parentName = fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, oneName)
						keynames = append(keynames, parentName)
						break
					}
				}
			}

		}

		if resourceSchema == nil {
			// error
		}

		if pValue, ok := value.(*[]byte); ok {
			pHash, err := app.hashAdapter.FromBytes(*pValue)
			if err != nil {
				return nil, err
			}

			sections := strings.Split(oneFieldName, groupNameDelimiterForTableNames)
			groupNames := sections[0 : len(sections)-1]
			parentName := strings.Join(groupNames, groupNameDelimiterForTableNames)
			retIns, err := app.retrieveResourceValue(
				resourceSchema,
				*pHash,
				parentName,
			)

			if err != nil {
				return nil, err
			}

			length := len(groupMethods)
			for i := 0; i < length; i++ {
				lastIndex := length - i - 1
				keyname := keynames[lastIndex]
				rootMethod := rootMethods[lastIndex]
				groupMethod := groupMethods[lastIndex]
				if builderIns, ok := app.builders[keyname]; ok {

					// initialize the builder:
					errorString := ""
					initializeName := rootMethod.Initialize()
					retValue, err := callMethodsOnInstance(
						[]string{
							initializeName,
						},
						builderIns,
						&errorString,
					)
					if err != nil {
						return nil, err
					}

					if errorString != "" {
						return nil, errors.New(errorString)
					}

					// add the value to the builder:
					errorString = ""
					fieldName := groupMethod.Element()
					retValue, err = callMethodOnInstanceWithParams(
						fieldName,
						retValue,
						&errorString,
						[]interface{}{
							retIns,
						},
					)
					if err != nil {
						str := fmt.Sprintf("there was an error while adding a value to the builder (keyname: %s), error: %s", keyname, err.Error())
						return nil, errors.New(str)
					}

					if errorString != "" {
						return nil, errors.New(errorString)
					}

					// trigger the builder:
					errorString = ""
					triggerName := rootMethod.Trigger()
					retValue, err = callMethodsOnInstance(
						[]string{
							triggerName,
						},
						retValue,
						&errorString,
					)
					if err != nil {
						return nil, err
					}

					if errorString != "" {
						return nil, errors.New(errorString)
					}

					// change the instance to compose:
					retIns = retValue
					continue
				}

				// error
			}

			return app.tokenBuilder.Create().
				CreatedOn(createdOn).
				WithContent(retIns.(tokens.Content)).
				Now()

		}

		// error
	}

	return nil, nil
}

func (app *resourceRepository) fetchRetrievalFields(
	root roots.Root,
) ([]string, []interface{}, error) {
	parentName := root.Name()
	chains := root.Chains()
	fieldNames, err := app.fetchRetrievalFieldsFromMethodChains(parentName, chains)
	if err != nil {
		return nil, nil, err
	}

	values := []interface{}{}
	for range fieldNames {
		var value []byte
		values = append(values, &value)
	}

	return fieldNames, values, nil
}

func (app *resourceRepository) fetchRetrievalFieldsFromGroup(
	parentName string,
	group groups.Group,
) ([]string, error) {
	name := group.Name()
	updatedParentName := fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, name)
	chains := group.Chains()
	return app.fetchRetrievalFieldsFromMethodChains(updatedParentName, chains)
}

func (app *resourceRepository) fetchRetrievalFieldsFromMethodChains(
	parentName string,
	chains groups.MethodChains,
) ([]string, error) {
	names := []string{}
	chainsList := chains.List()
	for _, oneChain := range chainsList {
		element := oneChain.Element()
		if element.IsGroup() {
			group := element.Group()
			subNames, err := app.fetchRetrievalFieldsFromGroup(parentName, group)
			if err != nil {
				return nil, err
			}

			names = append(names, subNames...)
		}

		if element.IsResource() {
			name := element.Resource().Name()
			fieldName := fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, name)
			names = append(names, fieldName)
		}
	}

	return names, nil
}

func (app *resourceRepository) retrieveResourceValue(
	resource schema_resources.Resource,
	hash hash.Hash,
	parentName string,
) (interface{}, error) {
	name := resource.Name()
	tableName := name
	if parentName != "" {
		tableName = fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, name)
	}

	propertyValues := []interface{}{}
	propertyNames := []string{}
	fieldsList := resource.Fields().List()
	for _, oneField := range fieldsList {
		typ := oneField.Type()
		propertyValue := app.generateValueFromType(typ)
		propertyValues = append(propertyValues, &propertyValue)

		propertyName := oneField.Name()
		propertyNames = append(propertyNames, propertyName)

	}

	propertyNamesStr := strings.Join(propertyNames, ",")
	queryStr := fmt.Sprintf("SELECT %s FROM %s WHERE hash = ?", propertyNamesStr, tableName)
	rows, err := app.dbPtr.Query(queryStr, hash.Bytes())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		str := fmt.Sprintf("the given hash (%s) do NOT match a %s instance", tableName, hash.String())
		return nil, errors.New(str)
	}

	err = rows.Scan(propertyValues...)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	resourceBuilderMethods := resource.Builder()
	if builderIns, ok := app.builders[tableName]; ok {
		// initialize the builder:
		errorString := ""
		initializeName := resourceBuilderMethods.Initialize()
		retValue, err := callMethodsOnInstance(
			[]string{
				initializeName,
			},
			builderIns,
			&errorString,
		)
		if err != nil {
			return nil, err
		}

		if errorString != "" {
			return nil, errors.New(errorString)
		}

		// set the builder instance:
		builderIns = retValue

		// pass the fields:
		for idx, oneField := range fieldsList {
			errorString := ""
			elementMethod := oneField.Methods().Element()
			pInterface := propertyValues[idx].(*interface{})
			retValue, err := callMethodOnInstanceWithParams(
				elementMethod,
				builderIns,
				&errorString,
				[]interface{}{
					*pInterface,
				},
			)
			if err != nil {
				return nil, err
			}

			if errorString != "" {
				return nil, errors.New(errorString)
			}

			// set the builder instance:
			builderIns = retValue
		}

		// trigger:
		errorString = ""
		triggerName := resourceBuilderMethods.Trigger()
		retValue, err = callMethodsOnInstance(
			[]string{
				triggerName,
			},
			builderIns,
			&errorString,
		)
		if err != nil {
			return nil, err
		}

		if errorString != "" {
			return nil, errors.New(errorString)
		}

		return retValue, nil
	}

	str := fmt.Sprintf("there is no resource builder for the provided tableName: %s", tableName)
	return nil, errors.New(str)
}

func (app *resourceRepository) generateValueFromType(typ field_types.Type) interface{} {
	if typ.IsKind() {
		pKind := typ.Kind()
		return app.generateValue(*pKind)
	}

	return app.generateValue(field_types.KindBytes)
}

func (app *resourceRepository) generateValue(kind uint8) interface{} {
	if kind == field_types.KindInteger {
		var value int
		return value
	}

	if kind == field_types.KindFloat {
		var value float64
		return value
	}

	if kind == field_types.KindString {
		var value string
		return value
	}

	var value []byte
	return value
}
