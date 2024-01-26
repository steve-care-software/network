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
	schema_groups "steve.care/network/domain/schemas/groups"
	schema_resources "steve.care/network/domain/schemas/groups/resources"
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

	_, err = app.txPtr.Exec("INSERT OR IGNORE INTO resource (hash, token, signature) VALUES (?, ?, ?)", ins.Hash().Bytes(), token.Hash().Bytes(), sigBytes)
	if err != nil {
		return err
	}

	return nil
}

func (app *resourceService) insertToken(ins tokens.Token) error {
	content := ins.Content()
	group := app.schema.Group()
	fkHash, fieldName, err := app.insertGroup(content, group, "")
	if err != nil {
		return err
	}

	queryStr := fmt.Sprintf("INSERT OR IGNORE INTO token (hash, %s, created_on) VALUES (?, ?, ?)", fieldName)
	_, err = app.txPtr.Exec(queryStr, ins.Hash().Bytes(), fkHash.Bytes(), ins.CreatedOn().Format(timeLayout))
	if err != nil {
		return err
	}

	return nil
}

func (app *resourceService) insertGroup(
	ins interface{},
	group schema_groups.Group,
	parentName string,
) (*hash.Hash, string, error) {
	name := group.Name()
	chains := group.Chains()
	updatedParentName := name
	if parentName != "" {
		updatedParentName = fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, name)
	}

	return app.insertChains(ins, chains, updatedParentName)
}

func (app *resourceService) insertChains(
	ins interface{},
	chains schema_groups.MethodChains,
	parentName string,
) (*hash.Hash, string, error) {
	list := chains.List()
	for _, oneChain := range list {
		pHash, fieldName, isInserted, err := app.insertChain(ins, oneChain, parentName)
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
) (*hash.Hash, string, bool, error) {
	errorString := ""
	conditionMethodName := chain.Condition()
	retValue, err := app.callMethodsOnInstance([]string{
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
		valueMethodName := chain.Value()
		retValue, err := app.callMethodsOnInstance([]string{
			valueMethodName,
		}, ins, &errorString)
		if err != nil {
			return nil, "", false, err
		}

		if errorString != "" {
			str := fmt.Sprintf("there was an error while calling a method chain's value method (name: %s) on a reference instance (type: %s), the error was: %s", valueMethodName, typeName, errorString)
			return nil, "", false, errors.New(str)
		}

		element := chain.Element()
		pHash, fieldName, err := app.insertElement(retValue, element, parentName)
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
) (*hash.Hash, string, error) {
	if element.IsResource() {
		resource := element.Resource()
		return app.insertResource(ins, resource, parentName)
	}

	group := element.Group()
	return app.insertGroup(ins, group, parentName)
}

func (app *resourceService) insertResource(
	ins interface{},
	resource schema_resources.Resource,
	parentName string,
) (*hash.Hash, string, error) {
	key := resource.Key()
	fieldNames := []string{
		key.Name(),
	}

	errorString := ""
	keyMethodNames := key.Methods()
	typeName := reflect.TypeOf(&ins).Elem().Name()
	retPkValue, err := app.callMethodsOnInstance(keyMethodNames, ins, &errorString)
	if err != nil {
		return nil, "", err
	}

	if errorString != "" {
		str := fmt.Sprintf("there was an error while calling a field key method (names: %s) on a reference instance (type: %s), the error was: %s", strings.Join(keyMethodNames, ","), typeName, errorString)
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
		methodNames := oneField.Methods()
		retValue, err := app.callMethodsOnInstance(methodNames, ins, &errorString)
		if err != nil {
			return nil, "", err
		}

		if errorString != "" {
			str := fmt.Sprintf("there was an error while calling a field method (names: %s) on a reference instance (type: %s), the error was: %s", strings.Join(methodNames, ","), typeName, errorString)
			return nil, "", errors.New(str)
		}

		fieldNames = append(fieldNames, oneField.Name())
		fieldValues = append(fieldValues, retValue)
		fieldValuePlaceHolders = append(fieldValuePlaceHolders, "?")
	}

	fieldNamesStr := strings.Join(fieldNames, ", ")
	fieldValuePlaceHoldersStr := strings.Join(fieldValuePlaceHolders, ", ")
	tableName := fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, resource.Name())
	queryStr := fmt.Sprintf("INSERT OR IGNORE INTO %s (%s) VALUES (%s)", tableName, fieldNamesStr, fieldValuePlaceHoldersStr)
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

	str := fmt.Sprintf("the returned value of the field key method (names: %s) on a reference instance (type: %s), was expected to contain []byte, but it was NOT", strings.Join(keyMethodNames, ","), typeName)
	return nil, "", errors.New(str)
}

func (app *resourceService) callMethodsOnInstance(
	methods []string,
	pInstance interface{},
	pErrorStr *string,
) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			value := fmt.Sprint(r)
			*pErrorStr = value
		}
	}()

	value := reflect.ValueOf(pInstance)
	for _, oneMethod := range methods {
		retValues := value.MethodByName(oneMethod).Call([]reflect.Value{})
		if len(retValues) != 1 {
			str := fmt.Sprintf("%d  values were returned, %d were expected, when calling the metod (name %s) in the method chain (%s)", len(retValues), 1, oneMethod, strings.Join(methods, ","))
			return nil, errors.New(str)
		}

		value = retValues[0]
	}

	return value.Interface(), nil
}

// Delete deletes a resource
func (app *resourceService) Delete(hash hash.Hash) error {
	return nil
}
