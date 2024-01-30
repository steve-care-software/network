package sqllites

import (
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"steve.care/network/applications"
	core_applications "steve.care/network/applications/applications"
	accounts_applications "steve.care/network/applications/applications/accounts"
	resources_applications "steve.care/network/applications/applications/resources"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/encryptors"
	"steve.care/network/domain/schemas"
	"steve.care/network/domain/schemas/groups"
	"steve.care/network/domain/schemas/groups/resources"
	"steve.care/network/domain/schemas/groups/resources/fields"
	field_types "steve.care/network/domain/schemas/groups/resources/fields/types"
)

type tableMetaData struct {
	resourceName string
	tableName    string
	fields       []string
	connections  []resources.Connection
}

type application struct {
	schema     schemas.Schema
	encryptor  encryptors.Encryptor
	adapter    accounts.Adapter
	bitrate    int
	basePath   string
	currentDb  *sql.DB
	currentTrx *sql.Tx
	baseSchema string
}

func createApplication(
	schema schemas.Schema,
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	bitrate int,
	basePath string,
	baseSchema string,
) applications.Application {
	out := application{
		schema:     schema,
		encryptor:  encryptor,
		adapter:    adapter,
		bitrate:    bitrate,
		basePath:   basePath,
		baseSchema: baseSchema,
		currentDb:  nil,
		currentTrx: nil,
	}

	return &out
}

// Init initializes the application
func (app *application) Init(name string) error {
	if app.isActive() {
		return errors.New(currentActiveErrorMsg)
	}

	err := app.openDatabaseIfNotAlready(name)
	if err != nil {
		return err
	}

	// init the schema:
	group := app.schema.Group()
	tableMetaDataList, err := app.initGroup("", group)
	if err != nil {
		return err
	}

	connectionsMap, err := app.initGroupForConnections(group, tableMetaDataList)
	if err != nil {
		return err
	}

	schema, err := app.generateSchema(tableMetaDataList, connectionsMap)
	if err != nil {
		return err
	}

	_, err = app.currentDb.Exec(schema)
	if err != nil {
		return err
	}

	// init the base schema:
	_, err = app.currentDb.Exec(app.baseSchema)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) generateSchema(metaData []*tableMetaData, connections map[string][]string) (string, error) {
	schema := ""
	tokenFieldNamesList := []string{}
	tokenFKNamesList := []string{}
	for _, oneMetaData := range metaData {
		tableName := oneMetaData.tableName
		foreignKeysStringList := []string{}
		if fkList, ok := connections[tableName]; ok {
			foreignKeysStringList = fkList
		}

		fieldsListString := ""
		fieldsAmount := len(oneMetaData.fields)
		for idx, oneField := range oneMetaData.fields {
			if (idx+1) == fieldsAmount && len(foreignKeysStringList) <= 0 {
				fieldsListString = fmt.Sprintf("%s\n%s\n", fieldsListString, oneField)
				continue
			}

			fieldsListString = fmt.Sprintf("%s\n%s%s", fieldsListString, oneField, ",")
		}

		foreignKeyString := ""
		if len(foreignKeysStringList) > 0 {
			pkAmount := len(foreignKeysStringList)
			for idx, oneForeignKey := range foreignKeysStringList {
				if (idx + 1) == pkAmount {
					foreignKeyString = fmt.Sprintf("%s\n%s\n", foreignKeyString, oneForeignKey)
					continue
				}

				foreignKeyString = fmt.Sprintf("%s\n%s%s", foreignKeyString, oneForeignKey, ",")
			}
		}

		// token field and FK:
		fieldName := fmt.Sprintf("%s %s,", tableName, "BLOB")
		fkName := fmt.Sprintf("FOREIGN KEY(%s) REFERENCES %s(hash)", tableName, tableName)
		tokenFKNamesList = append(tokenFKNamesList, fkName)
		tokenFieldNamesList = append(tokenFieldNamesList, fieldName)

		// drop and create tables:
		dropTableStr := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
		createTableStr := fmt.Sprintf("CREATE TABLE %s (%s%s);", tableName, fieldsListString, foreignKeyString)
		schema = fmt.Sprintf("%s\n\n%s\n%s", schema, dropTableStr, createTableStr)
	}

	dropTokenTable := "DROP TABLE IF EXISTS token;"
	createTokenTable := fmt.Sprintf(
		`CREATE TABLE token (
			hash BLOB PRIMARY KEY,
			%s
			created_on TEXT,
			%s
		);
	`,
		strings.Join(tokenFieldNamesList, ""),
		strings.Join(tokenFKNamesList, ","),
	)

	dropResourceTable := "DROP TABLE IF EXISTS resource;"
	createResourceTable := `CREATE TABLE resource (
			hash BLOB PRIMARY KEY,
			token BLOB,
			signature BLOB,
			FOREIGN KEY(token) REFERENCES token(hash)
		);
	`

	schema = fmt.Sprintf("%s\n\n%s\n%s\n\n%s\n%s", schema, dropTokenTable, createTokenTable, dropResourceTable, createResourceTable)
	return schema, nil
}

func (app *application) initGroupForConnections(group groups.Group, metaData []*tableMetaData) (map[string][]string, error) {
	chains := group.Chains()
	return app.initMethodChainsForConnections(chains, metaData)
}

func (app *application) initMethodChainsForConnections(chains groups.MethodChains, metaData []*tableMetaData) (map[string][]string, error) {
	list := chains.List()
	output := map[string][]string{}
	for _, oneChain := range list {
		retMap, err := app.initMethodChainForConnections(oneChain, metaData)
		if err != nil {
			return nil, err
		}

		for keyname, list := range retMap {
			output[keyname] = list
		}
	}

	return output, nil
}

func (app *application) initMethodChainForConnections(chain groups.MethodChain, metaData []*tableMetaData) (map[string][]string, error) {
	element := chain.Element()
	return app.initElementForConnections(element, metaData)
}

func (app *application) initElementForConnections(element groups.Element, metaData []*tableMetaData) (map[string][]string, error) {
	if element.IsGroup() {
		group := element.Group()
		return app.initGroupForConnections(group, metaData)
	}

	resource := element.Resource()
	tableName, list, err := app.initResourceForConnections(resource, metaData)
	if err != nil {
		return nil, err
	}

	return map[string][]string{
		tableName: list,
	}, nil
}

func (app *application) initResourceForConnections(resource resources.Resource, metaData []*tableMetaData) (string, []string, error) {
	resName := resource.Name()
	tableName, err := app.getTableNameByResourceName(resName, metaData)
	if err != nil {
		return "", nil, err
	}

	if !resource.HasConnections() {
		return tableName, []string{}, nil
	}

	connections := resource.Connections()
	list, err := app.initConnections(connections, metaData)
	if err != nil {
		return "", nil, err
	}

	return tableName, list, nil
}

func (app *application) initConnections(connections resources.Connections, metaData []*tableMetaData) ([]string, error) {
	output := []string{}
	list := connections.List()
	for _, oneConnection := range list {
		str, err := app.initConnection(oneConnection, metaData)
		if err != nil {
			return nil, err
		}

		output = append(output, str)
	}

	return output, nil
}

func (app *application) initConnection(connection resources.Connection, metaData []*tableMetaData) (string, error) {
	field := connection.Field()
	reference := connection.Reference()
	refResourceName := reference.Resource().Name()
	refTableName, err := app.getTableNameByResourceName(refResourceName, metaData)
	if err != nil {
		return "", err
	}

	refFieldName := reference.Field()
	return fmt.Sprintf("FOREIGN KEY(%s) REFERENCES %s(%s)", field, refTableName, refFieldName), nil
}

func (app *application) getTableNameByResourceName(resourceName string, metaData []*tableMetaData) (string, error) {
	for _, oneMetaData := range metaData {
		if oneMetaData.resourceName != resourceName {
			continue
		}

		return oneMetaData.tableName, nil
	}

	str := fmt.Sprintf("there is no resource named '%s' in the provided schema", resourceName)
	return "", errors.New(str)
}

func (app *application) initGroup(previousGroupName string, group groups.Group) ([]*tableMetaData, error) {
	name := group.Name()
	if previousGroupName != "" {
		name = fmt.Sprintf("%s%s%s", previousGroupName, groupNameDelimiterForTableNames, name)
	}

	chains := group.Chains()
	return app.initChains(name, chains)
}

func (app *application) initChains(groupName string, chains groups.MethodChains) ([]*tableMetaData, error) {
	output := []*tableMetaData{}
	list := chains.List()
	for _, oneChain := range list {
		retList, err := app.initChain(groupName, oneChain)
		if err != nil {
			return nil, err
		}

		output = append(output, retList...)
	}

	return output, nil
}

func (app *application) initChain(groupName string, chain groups.MethodChain) ([]*tableMetaData, error) {
	element := chain.Element()
	return app.initElement(groupName, element)
}

func (app *application) initElement(groupName string, element groups.Element) ([]*tableMetaData, error) {
	if element.IsGroup() {
		group := element.Group()
		return app.initGroup(groupName, group)
	}

	resource := element.Resource()
	pMetaData, err := app.initResource(groupName, resource)
	if err != nil {
		return nil, err
	}

	return []*tableMetaData{
		pMetaData,
	}, nil
}

func (app *application) initResource(groupName string, resource resources.Resource) (*tableMetaData, error) {
	key := resource.Key()
	keyFieldString, err := app.getTableFieldString(key)
	if err != nil {
		return nil, err
	}

	keyFieldStringWithPK := fmt.Sprintf("%s %s", keyFieldString, "PRIMARY KEY")
	fields := resource.Fields()
	fieldStringList, err := app.getTableFieldsStringList(fields)
	if err != nil {
		return nil, err
	}

	connections := []resources.Connection{}
	if resource.HasConnections() {
		connections = resource.Connections().List()
	}

	name := resource.Name()
	tableName := fmt.Sprintf("%s%s%s", groupName, groupNameDelimiterForTableNames, name)

	allFields := []string{
		keyFieldStringWithPK,
	}

	allFields = append(allFields, fieldStringList...)
	return &tableMetaData{
		resourceName: name,
		tableName:    tableName,
		fields:       allFields,
		connections:  connections,
	}, nil
}

func (app *application) getTableFieldsStringList(fields fields.Fields) ([]string, error) {
	output := []string{}
	list := fields.List()
	for _, oneField := range list {
		str, err := app.getTableFieldString(oneField)
		if err != nil {
			return nil, err
		}

		output = append(output, str)
	}

	return output, nil
}

func (app *application) getTableFieldString(field fields.Field) (string, error) {
	notNullString := " NOT NULL"
	if field.CanBeNil() {
		notNullString = ""
	}

	name := field.Name()
	typ := field.Type()
	kindString, err := app.getTableKindFromType(typ)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s%s", name, kindString, notNullString), nil
}

func (app *application) getTableKindFromType(typ field_types.Type) (string, error) {
	if typ.IsKind() {
		pKind := typ.Kind()
		return app.getTableKindString(*pKind)
	}

	return app.getTableKindString(field_types.KindBytes)
}

func (app *application) getTableKindString(kind uint8) (string, error) {
	if field_types.KindInteger == kind {
		return "INTEGER", nil
	}

	if field_types.KindFloat == kind {
		return "REAL", nil
	}

	if field_types.KindString == kind {
		return "TEXT", nil
	}

	return "BLOB", nil
}

// Begin begins the application
func (app *application) Begin(name string) (core_applications.Application, error) {
	err := app.openDatabaseIfNotAlready(name)
	if err != nil {
		return nil, err
	}

	return app.begin()
}

// Commit commits the application
func (app *application) Commit() error {
	err := app.currentTrx.Commit()
	if err != nil {
		return err
	}

	app.currentTrx = nil
	return nil
}

// Rollback rollbacks the application
func (app *application) Rollback() error {
	if !app.isActive() {
		return errors.New(notActiveErrorMsg)
	}

	err := app.currentTrx.Rollback()
	if err != nil {
		return err
	}

	app.currentTrx = nil
	return nil
}

// Close closes the application
func (app *application) Close() error {
	if !app.isActive() {
		return errors.New(notActiveErrorMsg)
	}

	err := app.currentDb.Close()
	if err != nil {
		return err
	}

	app.currentDb = nil
	app.currentTrx = nil
	return nil
}

func (app *application) begin() (core_applications.Application, error) {
	if !app.isTransactionActive() {
		trxApp, err := app.currentDb.Begin()
		if err != nil {
			return nil, err
		}

		app.currentTrx = trxApp
	}

	accountRepository := NewAccountRepository(
		app.encryptor,
		app.adapter,
		app.currentDb,
	)

	accountService := NewAccountService(
		accountRepository,
		app.encryptor,
		app.adapter,
		app.bitrate,
		app.currentTrx,
	)

	accountApplication := accounts_applications.NewApplication(
		accountRepository,
		accountService,
		app.bitrate,
	)

	resourceRepository := NewResourceRepository(
		app.schema,
		app.currentDb,
	)

	resourceService := NewResourceService(
		app.schema,
		app.currentTrx,
	)

	resourceApplication := resources_applications.NewApplication(
		resourceRepository,
		resourceService,
	)

	return core_applications.NewApplication(
		accountApplication,
		resourceApplication,
	), nil
}

func (app *application) openDatabaseIfNotAlready(name string) error {
	if !app.isDatabaseOpen() {
		database, err := app.open(name)
		if err != nil {
			return err
		}

		app.currentDb = database
	}

	return nil
}

func (app *application) isActive() bool {
	return app.isDatabaseOpen() &&
		app.currentDb != nil
}

func (app *application) isTransactionActive() bool {
	return app.currentTrx != nil
}

func (app *application) isDatabaseOpen() bool {
	return app.currentDb != nil
}

func (app *application) open(name string) (*sql.DB, error) {
	basePath := filepath.Join(app.basePath, name)
	return sql.Open("sqlite3", basePath)
}
