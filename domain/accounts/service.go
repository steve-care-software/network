package accounts

import (
	_ "github.com/mattn/go-sqlite3"
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases/transactions"
	"steve.care/network/domain/encryptors"
)

type service struct {
	encryptor        encryptors.Encryptor
	builder          Builder
	repository       Repository
	adapter          Adapter
	encryptorBuilder account_encryptors.Builder
	signerFactory    signers.Factory
	trx              transactions.Transaction
	bitrate          int
}

func createService(
	encryptor encryptors.Encryptor,
	builder Builder,
	repository Repository,
	adapter Adapter,
	encryptorBuilder account_encryptors.Builder,
	signerFactory signers.Factory,
	trx transactions.Transaction,
	bitrate int,
) Service {
	out := service{
		encryptor:        encryptor,
		trx:              trx,
		builder:          builder,
		repository:       repository,
		adapter:          adapter,
		encryptorBuilder: encryptorBuilder,
		signerFactory:    signerFactory,
		bitrate:          bitrate,
	}

	return &out
}

// Insert inserts an account
func (app *service) Insert(account Account, password []byte) error {
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

	return app.trx.Execute("INSERT INTO accounts (username, cipher) VALUES (?, ?)", username, cipher)
}

// Update updates an account
func (app *service) Update(credentials credentials.Credentials, criteria UpdateCriteria) error {
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

	return app.trx.Execute("UPDATE accounts set username = ?, cipher = ? where username = ?", updatedAccount.Username(), cipher, originUsername)
}

// Delete deletes an account
func (app *service) Delete(credentials credentials.Credentials) error {
	return app.trx.Execute("DELETE FROM accounts where username = ?", credentials.Username())
}
