package sqllites

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	"steve.care/network/domain/schemas"
	"steve.care/network/domain/schemas/roots"
	schema_groups "steve.care/network/domain/schemas/roots/groups"
	schema_resources "steve.care/network/domain/schemas/roots/groups/resources"
)

type resourceService struct {
	hashAdapter hash.Adapter
	schema      schemas.Schema
	txPtr       *sql.Tx
}

func createResourceService(
	hashAdapter hash.Adapter,
	schema schemas.Schema,
	txPtr *sql.Tx,
) resources.Service {
	out := resourceService{
		hashAdapter: hashAdapter,
		schema:      schema,
		txPtr:       txPtr,
	}

	return &out
}

// Insert inserts a resource
func (app *resourceService) Insert(ins resources.Resource) error {
	token := ins.Token()
	err := app.insertToken(token)
	if err != nil {
		return err
	}

	sigBytes, err := ins.Signature().Bytes()
	if err != nil {
		return err
	}

	_, err = app.txPtr.Exec("INSERT INTO resource (hash, token, signature) VALUES (?, ?, ?)", ins.Hash().Bytes(), token.Hash().Bytes(), sigBytes)
	if err != nil {
		return err
	}

	return nil
}

func (app *resourceService) insertToken(ins tokens.Token) error {
	content := ins.Content()
	root := app.schema.Root()
	fkHash, fieldName, err := app.insertRoot(content, root)
	if err != nil {
		return err
	}

	queryStr := fmt.Sprintf("INSERT INTO token (hash, %s, created_on) VALUES (?, ?, ?)", fieldName)
	_, err = app.txPtr.Exec(queryStr, ins.Hash().Bytes(), fkHash.Bytes(), ins.CreatedOn().Format(timeLayout))
	if err != nil {
		return err
	}

	return nil
}

func (app *resourceService) insertRoot(
	ins interface{},
	root roots.Root,
) (*hash.Hash, string, error) {
	name := root.Name()
	chains := root.Chains()
	return app.insertChains(ins, chains, name, root)
}

func (app *resourceService) insertGroup(
	ins interface{},
	group schema_groups.Group,
	parentName string,
	root roots.Root,
) (*hash.Hash, string, error) {
	name := group.Name()
	chains := group.Chains()
	updatedParentName := name
	if parentName != "" {
		updatedParentName = fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, name)
	}

	return app.insertChains(ins, chains, updatedParentName, root)
}

func (app *resourceService) insertChains(
	ins interface{},
	chains schema_groups.MethodChains,
	parentName string,
	root roots.Root,
) (*hash.Hash, string, error) {
	list := chains.List()
	for _, oneChain := range list {
		pHash, fieldName, isInserted, err := app.insertChain(ins, oneChain, parentName, root)
		if err != nil {
			return nil, "", err
		}

		if !isInserted {
			continue
		}

		return pHash, fieldName, nil
	}

	return nil, "", nil
}

func (app *resourceService) insertChain(
	ins interface{},
	chain schema_groups.MethodChain,
	parentName string,
	root roots.Root,
) (*hash.Hash, string, bool, error) {

	errorString := ""
	conditionMethodName := chain.Condition()
	retValue, err := callMethodsOnInstance([]string{
		conditionMethodName,
	}, ins, &errorString)
	if err != nil {
		return nil, "", false, err
	}

	typeName := reflect.TypeOf(&ins).Elem().Name()
	if errorString != "" {
		str := fmt.Sprintf("there was an error while calling a method chain's condition method (name: %s) on a reference instance (type: %s), the error was: %s", conditionMethodName, typeName, errorString)
		return nil, "", false, errors.New(str)
	}

	if boolValue, ok := retValue.(bool); ok {
		if !boolValue {
			return nil, "", false, nil
		}

		errorString := ""
		retrieverMethodNames := chain.Retriever()
		retValue, err := callMethodsOnInstance(retrieverMethodNames, ins, &errorString)
		if err != nil {
			return nil, "", false, err
		}

		if errorString != "" {
			str := fmt.Sprintf("there was an error while calling a method chain's value method (chain: %s) on a reference instance (type: %s), the error was: %s", strings.Join(retrieverMethodNames, ","), typeName, errorString)
			return nil, "", false, errors.New(str)
		}

		element := chain.Element()
		pHash, fieldName, err := app.insertElement(retValue, element, parentName, root)
		if err != nil {
			return nil, "", false, err
		}

		return pHash, fieldName, true, nil
	}

	str := fmt.Sprintf("there was an error while calling a method chain's condition method (name: %s) on a reference instance (type: %s), the returned value was expected to be a bool, but it was NOT", conditionMethodName, typeName)
	return nil, "", false, errors.New(str)
}

func (app *resourceService) insertElement(
	ins interface{},
	element schema_groups.Element,
	parentName string,
	root roots.Root,
) (*hash.Hash, string, error) {
	if element.IsResource() {
		resource := element.Resource()
		return app.insertResource(ins, resource, parentName, root)
	}

	group := element.Group()
	return app.insertGroup(ins, group, parentName, root)
}

func (app *resourceService) insertResource(
	ins interface{},
	resource schema_resources.Resource,
	parentName string,
	root roots.Root,
) (*hash.Hash, string, error) {
	key := resource.Key()
	fieldNames := []string{
		key.Name(),
	}

	errorString := ""
	keyMethods := key.Methods()
	typeName := reflect.TypeOf(&ins).Elem().Name()
	retPkValue, err := callMethodsOnInstance(keyMethods.Retriever(), ins, &errorString)
	if err != nil {
		return nil, "", err
	}

	if errorString != "" {
		str := fmt.Sprintf("there was an error while calling a field key method (names: %s) on a reference instance (type: %s), the error was: %s", strings.Join(keyMethods.Retriever(), ","), typeName, errorString)
		return nil, "", errors.New(str)
	}

	fieldValues := []interface{}{
		retPkValue,
	}

	fieldValuePlaceHolders := []string{
		"?",
	}

	fieldsList := resource.Fields().List()
	for _, oneField := range fieldsList {
		errorString := ""
		methods := oneField.Methods()
		retValue, err := callMethodsOnInstance(methods.Retriever(), ins, &errorString)
		if err != nil {
			return nil, "", err
		}

		if errorString != "" {
			str := fmt.Sprintf("there was an error while calling a field method (names: %s) on a reference instance (type: %s), the error was: %s", strings.Join(methods.Retriever(), ","), typeName, errorString)
			return nil, "", errors.New(str)
		}

		if retValue == nil && !oneField.CanBeNil() {
			str := fmt.Sprintf("the field (resource: %s, name: %s) is nil but is set as 'cannot be nil' in the schema", resource.Name(), oneField.Name())
			return nil, "", errors.New(str)
		}

		// do not set the field in the query if the value is nil:
		if retValue == nil {
			continue
		}

		typ := oneField.Type()
		if typ.IsDependency() {
			// call the depenency retriever on the instance:
			dependency := typ.Dependency()
			retriever := dependency.Retriever()
			errorString := ""
			retValue, err := callMethodsOnInstance([]string{retriever}, ins, &errorString)
			if err != nil {
				return nil, "", err
			}

			if errorString != "" {
				str := fmt.Sprintf("there was an error while calling a field dependency method (name: %s) on a reference instance (type: %s), the error was: %s", retriever, typeName, errorString)
				return nil, "", errors.New(str)
			}

			// fetch the resource:
			groupNames := dependency.Groups()
			resourceName := dependency.Resource()
			path := append(groupNames, resourceName)
			retResourceSchema, err := root.Search(path)
			if err != nil {
				return nil, "", err
			}

			// generate the parent name:
			dependencyParentName := strings.Join(groupNames, groupNameDelimiterForTableNames)

			// insert the resource value:
			_, _, err = app.insertResource(retValue, retResourceSchema, dependencyParentName, root)
			if err != nil {
				return nil, "", err
			}
		}

		fieldNames = append(fieldNames, oneField.Name())
		fieldValues = append(fieldValues, retValue)
		fieldValuePlaceHolders = append(fieldValuePlaceHolders, "?")
	}

	fieldNamesStr := strings.Join(fieldNames, ", ")
	fieldValuePlaceHoldersStr := strings.Join(fieldValuePlaceHolders, ", ")
	tableName := fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, resource.Name())
	queryStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, fieldNamesStr, fieldValuePlaceHoldersStr)
	_, err = app.txPtr.Exec(queryStr, fieldValues...)
	if err != nil {
		return nil, "", err
	}

	if casted, ok := retPkValue.([]byte); ok {
		pHash, err := app.hashAdapter.FromBytes(casted)
		if err != nil {
			return nil, "", err
		}

		return pHash, tableName, nil
	}

	str := fmt.Sprintf("the returned value of the field key method (names: %s) on a reference instance (type: %s), was expected to contain []byte, but it was NOT", strings.Join(keyMethods.Retriever(), ","), typeName)
	return nil, "", errors.New(str)
}

// Delete deletes a resource
func (app *resourceService) Delete(hash hash.Hash) error {
	return nil
}
