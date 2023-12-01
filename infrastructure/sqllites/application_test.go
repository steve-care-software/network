package sqllites

import (
	"reflect"
	"testing"

	"steve.care/network/applications"
	"steve.care/network/domain/credentials"
	"steve.care/network/infrastructure/edwards25519"
	"steve.care/network/infrastructure/jsons"
)

func TestApplication_Account_InsertThenRetrieve_Success(t *testing.T) {
	appIns := applications.NewApplication(
		NewApplication("./"),
		edwards25519.NewEncryptor(),
		jsons.NewAccountAdapter(),
		4096,
	)

	name := "firstExec"
	defer appIns.Close(name)

	// init wit our schema:
	schema := getSchema()
	firstAppIns, err := appIns.InitInMemory(name, schema)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	accountAppIns, err := firstAppIns.Accounts()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	username := "rogerCyr"
	password := []byte("this is my password")
	credentials, err := credentials.NewBuilder().Create().WithUsername(username).WithPassword(password).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = accountAppIns.Insert(credentials)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = appIns.Commit(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retAccount, err := accountAppIns.Retrieve(credentials)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(username, retAccount.Username()) {
		t.Errorf("the returned account is invalid")
		return
	}

}
