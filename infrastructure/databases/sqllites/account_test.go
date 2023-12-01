package sqllites

import (
	"reflect"
	"testing"

	"steve.care/network/domain/accounts"
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
	"steve.care/network/infrastructure/encryptors/edwards25519"
	"steve.care/network/infrastructure/jsons"
)

func TestAccount_InsertThenRetrieve_Success(t *testing.T) {
	bitrate := 4096
	encryptor, err := encryptors.NewBuilder().Create().
		WithBitRate(bitrate).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	signer := signers.NewFactory().Create()
	username := "rogerCyr"
	accountIns, err := accounts.NewBuilder().Create().
		WithUsername(username).
		WithEncryptor(encryptor).
		WithSigner(signer).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	schema := GetSchema()
	dbPtr, err := openThenPrepareSQLInMemoryForTests(schema)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// defer close
	defer dbPtr.Close()

	repository := NewAccountRepository(
		edwards25519.NewApplication(),
		jsons.NewAccountAdapter(),
		dbPtr,
	)

	service := NewAccountService(
		edwards25519.NewApplication(),
		repository,
		jsons.NewAccountAdapter(),
		4096,
		dbPtr,
	)

	password := []byte("this is my passwprd")
	err = service.Insert(accountIns, password)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	credentials, err := credentials.NewBuilder().Create().WithUsername(username).WithPassword(password).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retAccount, err := repository.Retrieve(credentials)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(accountIns, retAccount) {
		t.Errorf("the returned account is invalid")
		return
	}

}
