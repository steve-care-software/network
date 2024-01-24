package sqllites

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"steve.care/network/domain/credentials"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
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
	appIns := NewApplication(
		NewSchemaFactory(
			keyFieldName,
		),
		edwards25519.NewEncryptor(),
		jsons.NewAccountAdapter(),
		4096,
		dbDir,
	)

	dbName := "testdb"
	defer func() {
		path := filepath.Join(dbDir, dbName)
		os.Remove(path)
	}()

	// close:
	defer appIns.Close()

	// init out app:
	err := appIns.Init()
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
