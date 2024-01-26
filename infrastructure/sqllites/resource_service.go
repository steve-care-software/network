package sqllites

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	tokens_dashboards "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/dashboards"
	"steve.care/network/domain/schemas"
	"steve.care/network/domain/schemas/groups"
	schema_resources "steve.care/network/domain/schemas/groups/resources"
)

type resourceService struct {
	schema schemas.Schema
	txPtr  *sql.Tx
}

func createResourceService(
	schema schemas.Schema,
	txPtr *sql.Tx,
) resources.Service {
	out := resourceService{
		schema: schema,
		txPtr:  txPtr,
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
	currentGroupName := "resources"
	content := ins.Content()
	group, err := app.schema.Groups().Fetch(currentGroupName)
	if err != nil {
		return err
	}

	if content.IsDashboard() {
		dashboard := content.Dashboard()
		fkHash, err := app.insertDashboard(dashboard, group, currentGroupName)
		if err != nil {
			return err
		}

		_, err = app.txPtr.Exec("INSERT OR IGNORE INTO token (hash, dashboards_viewport, created_on) VALUES (?, ?, ?)", ins.Hash().Bytes(), fkHash.Bytes(), ins.CreatedOn().Format(timeLayout))
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (app *resourceService) insertDashboard(ins tokens_dashboards.Dashboard, group groups.Group, parentName string) (*hash.Hash, error) {
	currentGroupName := "dashboards"
	group, err := group.Elements().Search("dashboards")
	if err != nil {
		return nil, err
	}

	concatGroupName := fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, currentGroupName)
	if ins.IsDashboard() {

	}

	if ins.IsWidget() {

	}

	viewport := ins.Viewport()
	return app.insertDashboardViewport(viewport, group, concatGroupName)
}

func (app *resourceService) insertDashboardViewport(
	ins viewports.Viewport,
	group groups.Group,
	parentName string,
) (*hash.Hash, error) {
	resource, err := group.Elements().Resource("viewport")
	if err != nil {
		return nil, err
	}

	err = app.insertResource(ins, resource, parentName)
	if err != nil {
		return nil, err
	}

	insHash := ins.Hash()
	return &insHash, nil
}

func (app *resourceService) insertResource(
	ins interface{},
	resource schema_resources.Resource,
	parentName string,
) error {
	key := resource.Key()
	fieldNames := []string{
		key.Name(),
	}

	errorString := ""
	keyMethodNames := key.Methods()
	typeName := reflect.TypeOf(&ins).Elem().Name()
	retPkValue, err := app.callMethodsOnInstance(keyMethodNames, ins, &errorString)
	if err != nil {
		return err
	}

	if errorString != "" {
		str := fmt.Sprintf("there was an error while calling a field key method (names: %s) on a reference instance (type: %s), the error was: %s", strings.Join(keyMethodNames, ","), typeName, errorString)
		return errors.New(str)
	}

	fieldValues := []interface{}{
		retPkValue,
	}

	fieldsList := resource.Fields().List()
	for _, oneField := range fieldsList {
		errorString := ""
		methodNames := oneField.Methods()
		retValue, err := app.callMethodsOnInstance(methodNames, ins, &errorString)
		if err != nil {
			return err
		}

		if errorString != "" {
			str := fmt.Sprintf("there was an error while calling a field method (names: %s) on a reference instance (type: %s), the error was: %s", strings.Join(methodNames, ","), typeName, errorString)
			return errors.New(str)
		}

		fieldNames = append(fieldNames, oneField.Name())
		fieldValues = append(fieldValues, retValue)
	}

	fieldNamesStr := strings.Join(fieldNames, ", ")
	tableName := fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, resource.Name())
	queryStr := fmt.Sprintf("INSERT OR IGNORE INTO %s (%s) VALUES (?, ?, ?)", tableName, fieldNamesStr)
	_, err = app.txPtr.Exec(queryStr, fieldValues...)
	if err != nil {
		return err
	}

	return nil
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
