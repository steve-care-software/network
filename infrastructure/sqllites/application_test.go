package sqllites

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	token_dashboards "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/dashboards"
	"steve.care/network/infrastructure/edwards25519"
	"steve.care/network/infrastructure/jsons"
)

type resourceExec struct {
	name     string
	resource resources.Resource
}

func TestApplication_Account_InsertThenRetrieve_Success(t *testing.T) {
	dbDir := "./test_files"
	keyFieldName := "hash"
	keyFieldMethodNames := []string{
		"Hash",
		"Bytes",
	}

	baseSchema := getSchema()
	schema, err := NewSchemaFactory(
		keyFieldName,
		keyFieldMethodNames,
	).Create()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	appIns := NewApplication(
		schema,
		edwards25519.NewEncryptor(),
		jsons.NewAccountAdapter(),
		4096,
		dbDir,
		baseSchema,
	)

	dbName := "testdb"
	defer func() {
		path := filepath.Join(dbDir, dbName)
		os.Remove(path)
	}()

	// close:
	defer appIns.Close()

	// init out app:
	err = appIns.Init(dbName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// begin our app:
	initAppIns, err := appIns.Begin(dbName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	accAppIns := initAppIns.Accounts()
	username := "rogerCyr"
	password := []byte("this is my password")
	credentials, err := credentials.NewBuilder().Create().WithUsername(username).WithPassword(password).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// insert account:
	err = accAppIns.Insert(credentials)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = appIns.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retAccount, err := accAppIns.Retrieve(credentials)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(username, retAccount.Username()) {
		t.Errorf("the returned account is invalid")
		return
	}

	// delete:
	/*err = accAppIns.Delete(credentials)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = appIns.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	exists, err := accAppIns.Exists(username)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if exists {
		t.Errorf("the account (username: %s) was NOT expected to exists", username)
		return
	}*/
}

func TestApplication_Resources_Success(t *testing.T) {
	dbDir := "./test_files"
	keyFieldName := "hash"
	keyFieldMethodNames := []string{
		"Hash",
		"Bytes",
	}

	baseSchema := getSchema()
	schema, err := NewSchemaFactory(
		keyFieldName,
		keyFieldMethodNames,
	).Create()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	appIns := NewApplication(
		schema,
		edwards25519.NewEncryptor(),
		jsons.NewAccountAdapter(),
		4096,
		dbDir,
		baseSchema,
	)

	dbName := "testdb"
	defer func() {
		path := filepath.Join(dbDir, dbName)
		os.Remove(path)
	}()

	// close:
	defer appIns.Close()

	// init out app:
	err = appIns.Init(dbName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// begin our app:
	initAppIns, err := appIns.Begin(dbName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the resources application:
	resAppIns := initAppIns.Resources()

	// build viewport resource:
	viewport := viewports.NewViewportForTests(uint(33), uint(45))
	token := tokens.NewTokenWithDashboardForTests(
		token_dashboards.NewDashboardWithViewportForTests(
			viewport,
		),
	)

	msg := token.Hash().Bytes()
	signature, err := signers.NewFactory().Create().Sign(msg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	resource := resources.NewResourceForTests(token, signature)

	// insert account:
	err = resAppIns.Insert(resource)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = appIns.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	//resHash := resource.Hash()
	resHash := resource.Hash()
	retResource, err := resAppIns.RetrieveByHash(resHash)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(retResource.Hash().Bytes(), resHash.Bytes()) {
		t.Errorf("the returned resource is invalid")
		return
	}
}
