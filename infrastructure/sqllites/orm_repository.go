package sqllites

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/orms"
	"steve.care/network/domain/orms/schemas/roots"
	"steve.care/network/domain/orms/schemas/roots/groups"
	schema_resources "steve.care/network/domain/orms/schemas/roots/groups/resources"
	field_types "steve.care/network/domain/orms/schemas/roots/groups/resources/fields/types"
	"steve.care/network/domain/orms/skeletons"
)

type ormRepository struct {
	hashAdapter hash.Adapter
	builders    map[string]interface{}
	skeleton    skeletons.Skeleton
	dbPtr       *sql.DB
}

func createOrmRepository(
	hashAdapter hash.Adapter,
	builders map[string]interface{},
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
) orms.Repository {
	out := ormRepository{
		hashAdapter: hashAdapter,
		builders:    builders,
		skeleton:    skeleton,
		dbPtr:       dbPtr,
	}

	return &out
}

// AmountByQuery returns the amount of instance by criteria
func (app *ormRepository) AmountByQuery(query hash.Hash) (uint, error) {
	return 0, nil
}

// ListByQuery lists insatnce hashes by criteria
func (app *ormRepository) ListByQuery(query hash.Hash) ([]hash.Hash, error) {
	return nil, nil
}

// RetrieveByQuery retrieves an instance by criteria
func (app *ormRepository) RetrieveByQuery(query hash.Hash) (orms.Instance, error) {
	return nil, nil
}

// RetrieveByHash retrieves an instance by hash
func (app *ormRepository) RetrieveByHash(path []string, hash hash.Hash) (orms.Instance, error) {
	return nil, nil
}

func (app *ormRepository) retrieveInstanceByRootsAndHash(
	hash hash.Hash,
	roots roots.Roots,
) (orms.Instance, error) {
	list := roots.List()
	for _, oneRoot := range list {
		retInstance, err := app.retrieveInstanceByRootAndHash(hash, oneRoot)
		if err != nil {
			continue
		}

		return retInstance, nil
	}

	str := fmt.Sprintf("the instance (hash: %s) could not be retrieved", hash.String())
	return nil, errors.New(str)
}

func (app *ormRepository) retrieveInstanceByRootAndHash(
	hash hash.Hash,
	root roots.Root,
) (orms.Instance, error) {
	/*fieldNames, values, err := app.fetchRetrievalFields(root)
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
		if pValue, ok := values[idx].(*[]byte); ok {
			if pValue == nil || len(*pValue) <= 0 {
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
			pHash, err := app.hashAdapter.FromBytes(*pValue)
			if err != nil {
				return nil, err
			}

			sections := strings.Split(oneFieldName, groupNameDelimiterForTableNames)
			groupNames := sections[0 : len(sections)-1]
			resourceParentName := strings.Join(groupNames, groupNameDelimiterForTableNames)
			retIns, err := app.retrieveResourceValue(
				resourceSchema,
				*pHash,
				resourceParentName,
				root,
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

					// only add the value to the builder element method when the value is NOT nil:
					if retIns != nil {
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
	}*/

	return nil, nil
}

func (app *ormRepository) fetchRetrievalFields(
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

func (app *ormRepository) fetchRetrievalFieldsFromGroup(
	parentName string,
	group groups.Group,
) ([]string, error) {
	name := group.Name()
	updatedParentName := fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, name)
	chains := group.Chains()
	return app.fetchRetrievalFieldsFromMethodChains(updatedParentName, chains)
}

func (app *ormRepository) fetchRetrievalFieldsFromMethodChains(
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

func (app *ormRepository) retrieveResourceValue(
	resource schema_resources.Resource,
	hash hash.Hash,
	parentName string,
	root roots.Root,
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
		propertyName := oneField.Name()

		propertyValues = append(propertyValues, &propertyValue)
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
		str := fmt.Sprintf("the given hash (%s) do NOT match a %s instance", hash.String(), tableName)
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

			if pCasted, ok := propertyValues[idx].(*interface{}); ok {
				// only call the builder element method when the value is NOT nil:
				if *pCasted != nil {
					casted := *pCasted
					typ := oneField.Type()
					if typ.IsDependency() {
						if castedBytes, ok := (casted.([]byte)); ok {
							// fetch the dependency resource:
							dependency := typ.Dependency()
							groupNames := dependency.Groups()
							resourceName := dependency.Resource()
							path := append(groupNames, resourceName)
							dependencyResourceSchema, err := root.Search(path)
							if err != nil {
								return nil, err
							}

							// create the parent name:
							dependencyParentName := strings.Join(groupNames, groupNameDelimiterForTableNames)

							// build an hash from the value:
							pHash, err := app.hashAdapter.FromBytes(castedBytes)
							if err != nil {
								return nil, err
							}

							// fetch the resource value:
							retResourceValue, err := app.retrieveResourceValue(
								dependencyResourceSchema,
								*pHash,
								dependencyParentName,
								root,
							)

							if err != nil {
								return nil, err
							}

							// use the resoruce value as the value in the builder:
							casted = retResourceValue
						}

					}

					errorString := ""
					elementMethod := oneField.Methods().Element()
					retValue, err := callMethodOnInstanceWithParams(
						elementMethod,
						builderIns,
						&errorString,
						[]interface{}{
							casted,
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

			}

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

func (app *ormRepository) generateValueFromType(typ field_types.Type) interface{} {
	if typ.IsKind() {
		pKind := typ.Kind()
		return app.generateValue(*pKind)
	}

	kind := typ.Dependency().Kind()
	return app.generateValue(kind)
}

func (app *ormRepository) generateValue(kind uint8) interface{} {
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
