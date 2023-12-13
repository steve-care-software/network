package sqllites

import (
	"reflect"
	"testing"

	"steve.care/network/domain/credentials"
	"steve.care/network/infrastructure/edwards25519"
	"steve.care/network/infrastructure/jsons"
)

func TestApplication_Account_InsertThenRetrieve_Success(t *testing.T) {
	appIns := NewApplication(
		edwards25519.NewEncryptor(),
		jsons.NewAccountAdapter(),
		4096,
		"./",
	)

	// close:
	defer appIns.Close()

	// init wit our schema:
	schema := getSchema()
	firstAppIns, err := appIns.InitInMemory(schema)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	firstAccAppIns, err := firstAppIns.Accounts()
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

	// insert:
	err = firstAccAppIns.Insert(credentials)
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

	retAccount, err := firstAccAppIns.Retrieve(credentials)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(username, retAccount.Username()) {
		t.Errorf("the returned account is invalid")
		return
	}

	// delete:
	secAppIns, err := appIns.BeginInMemory()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	secAccountAppIns, err := secAppIns.Accounts()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = secAccountAppIns.Delete(credentials)
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

	exists, err := secAccountAppIns.Exists(username)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if exists {
		t.Errorf("the account (username: %s) was NOT expected to exists", username)
		return
	}
}
