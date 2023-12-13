package sqllites

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"steve.care/network/domain/accounts"
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/encryptors"
)

type accountService struct {
	encryptor        encryptors.Encryptor
	builder          accounts.Builder
	repository       accounts.Repository
	adapter          accounts.Adapter
	encryptorBuilder account_encryptors.Builder
	signerFactory    signers.Factory
	bitrate          int
	txPtr            *sql.Tx
}

func createAccountService(
	encryptor encryptors.Encryptor,
	builder accounts.Builder,
	repository accounts.Repository,
	adapter accounts.Adapter,
	encryptorBuilder account_encryptors.Builder,
	signerFactory signers.Factory,
	bitrate int,
	txPtr *sql.Tx,
) accounts.Service {
	out := accountService{
		encryptor:        encryptor,
		builder:          builder,
		repository:       repository,
		adapter:          adapter,
		encryptorBuilder: encryptorBuilder,
		signerFactory:    signerFactory,
		bitrate:          bitrate,
		txPtr:            txPtr,
	}

	return &out
}

// Insert inserts an account
func (app *accountService) Insert(account accounts.Account, password []byte) error {
	bytes, err := app.adapter.ToBytes(account)
	if err != nil {
		return err
	}

	cipher, err := app.encryptor.Encrypt(bytes, password)
	if err != nil {
		return err
	}

	username := account.Username()
	if err != nil {
		return err
	}

	_, err = app.txPtr.Exec("INSERT INTO accounts (username, cipher) VALUES (?, ?)", username, cipher)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

// Update updates an account
func (app *accountService) Update(credentials credentials.Credentials, criteria accounts.UpdateCriteria) error {
	originAccount, err := app.repository.Retrieve(credentials)
	if err != nil {
		return err
	}

	originUsername := originAccount.Username()
	originSigner := originAccount.Signer()
	originEncryptor := originAccount.Encryptor()

	builder := app.builder.Create().
		WithUsername(originUsername).
		WithEncryptor(originEncryptor).
		WithSigner(originSigner)

	if criteria.ChangeEncryptor() {
		encryptor, err := app.encryptorBuilder.Create().
			WithBitRate(app.bitrate).
			Now()

		if err != nil {
			return err
		}

		builder.WithEncryptor(encryptor)
	}

	if criteria.ChangeSigner() {
		signer := app.signerFactory.Create()
		builder.WithSigner(signer)
	}

	if criteria.HasUsername() {
		username := criteria.Username()
		builder.WithUsername(username)
	}

	updatedAccount, err := builder.Now()
	if err != nil {
		return err
	}

	bytes, err := app.adapter.ToBytes(updatedAccount)
	if err != nil {
		return err
	}

	updatedPassword := credentials.Password()
	if criteria.HasPassword() {
		updatedPassword = criteria.Password()
	}

	cipher, err := app.encryptor.Encrypt(bytes, updatedPassword)
	if err != nil {
		return err
	}

	_, err = app.txPtr.Exec("UPDATE accounts set username = ?, cipher = ? where username = ?", updatedAccount.Username(), cipher, originUsername)
	if err != nil {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}

// Delete deletes an account
func (app *accountService) Delete(credentials credentials.Credentials) error {
	username := credentials.Username()
	_, err := app.repository.Retrieve(credentials)
	if err == nil {
		str := fmt.Sprintf("there is no account related to the provided credentials (username: %s)", username)
		return errors.New(str)
	}

	_, err = app.txPtr.Exec("DELETE FROM accounts WHERE username = ?", username)
	if err != nil {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}
