package sqllites

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	encryptor_applications "steve.care/network/applications/applications/encryptors"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
)

type accountService struct {
	encryptorApp     encryptor_applications.Application
	builder          accounts.Builder
	repository       accounts.Repository
	adapter          accounts.Adapter
	encryptorBuilder encryptors.Builder
	signerFactory    signers.Factory
	bitrate          int
	dbPtr            *sql.DB
}

func createAccountService(
	encryptorApp encryptor_applications.Application,
	builder accounts.Builder,
	repository accounts.Repository,
	adapter accounts.Adapter,
	encryptorBuilder encryptors.Builder,
	signerFactory signers.Factory,
	bitrate int,
	dbPtr *sql.DB,
) accounts.Service {
	out := accountService{
		encryptorApp:     encryptorApp,
		builder:          builder,
		repository:       repository,
		adapter:          adapter,
		encryptorBuilder: encryptorBuilder,
		signerFactory:    signerFactory,
		bitrate:          bitrate,
		dbPtr:            dbPtr,
	}

	return &out
}

// Insert inserts an account
func (app *accountService) Insert(account accounts.Account, password []byte) error {
	bytes, err := app.adapter.ToBytes(account)
	if err != nil {
		return err
	}

	cipher, err := app.encryptorApp.Encrypt(bytes, password)
	if err != nil {
		return err
	}

	// insert
	username := account.Username()
	stmt, err := app.dbPtr.Prepare("INSERT INTO accounts (username, cipher) VALUES (?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()
	res, err := stmt.Exec(username, cipher)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		str := fmt.Sprintf("the account could not be deleted properly, %d rows affected were expected, %d were in reality affected", 1, affected)
		return errors.New(str)
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

	cipher, err := app.encryptorApp.Encrypt(bytes, updatedPassword)
	if err != nil {
		return err
	}

	stmt, err := app.dbPtr.Prepare("UPDATE accounts set username = ?, cipher = ? where username = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	res, err := stmt.Exec(updatedAccount.Username(), cipher, originUsername)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		str := fmt.Sprintf("the account could not be deleted properly, %d rows affected were expected, %d were in reality affected", 1, affected)
		return errors.New(str)
	}

	return nil
}

// Delete deletes an account
func (app *accountService) Delete(credentials credentials.Credentials) error {
	stmt, err := app.dbPtr.Prepare("DELETE FROM accounts where username = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	username := credentials.Username()
	res, err := stmt.Exec(username)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		str := fmt.Sprintf("the account could not be deleted properly, %d rows affected were expected, %d were in reality affected", 1, affected)
		return errors.New(str)
	}

	return nil
}
