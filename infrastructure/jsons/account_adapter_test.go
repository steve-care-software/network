package jsons

import (
	"reflect"
	"testing"

	"steve.care/network/domain/accounts"
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
)

func TestAccountAdapter_Success(t *testing.T) {
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

	accountAdapter := NewAccountAdapter()
	js, err := accountAdapter.ToBytes(accountIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := accountAdapter.ToInstance(js)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(accountIns, retIns) {
		t.Errorf("the returned account instance is invalid")
		return
	}
}
