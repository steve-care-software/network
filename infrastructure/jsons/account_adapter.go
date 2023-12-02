package jsons

import (
	"encoding/json"

	"steve.care/network/domain/accounts"
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	json_accounts "steve.care/network/infrastructure/jsons/accounts"
)

type accountAdapter struct {
	builder          accounts.Builder
	encryptorAdapter encryptors.Adapter
	signerAdapter    signers.Adapter
}

func createAccountAdapter(
	builder accounts.Builder,
	encryptorAdapter encryptors.Adapter,
	signerAdapter signers.Adapter,
) accounts.Adapter {
	out := accountAdapter{
		builder:          builder,
		encryptorAdapter: encryptorAdapter,
		signerAdapter:    signerAdapter,
	}

	return &out
}

// ToBytes converts an account to bytes
func (app *accountAdapter) ToBytes(ins accounts.Account) ([]byte, error) {
	username := ins.Username()
	signerBytes, err := ins.Signer().Bytes()
	if err != nil {
		return nil, err
	}

	encryptor := ins.Encryptor()
	encryptorBytes := app.encryptorAdapter.ToBytes(encryptor)
	insAccount := json_accounts.Account{
		Username:  username,
		Encryptor: encryptorBytes,
		Signer:    signerBytes,
	}

	return json.Marshal(insAccount)
}

// ToInstance converts bytes to account
func (app *accountAdapter) ToInstance(bytes []byte) (accounts.Account, error) {
	ptr := new(json_accounts.Account)
	err := json.Unmarshal(bytes, ptr)
	if err != nil {
		return nil, err
	}

	signer, err := app.signerAdapter.ToSigner(ptr.Signer)
	if err != nil {
		return nil, err
	}

	encryptor, err := app.encryptorAdapter.ToEncryptor(ptr.Encryptor)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithUsername(ptr.Username).
		WithEncryptor(encryptor).
		WithSigner(signer).
		Now()
}
