package sqllites

import (
	"database/sql"

	encryptor_applications "steve.care/network/applications/applications/encryptors"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
)

// GetSchema returns the sql schema
func GetSchema() string {
	return `
		DROP TABLE IF EXISTS accounts;
		CREATE TABLE accounts(username TEXT PRIMARY KEY, cipher BLOB);
	`
}

// NewAccountRepository creates a new account reposiotry
func NewAccountRepository(
	encryptorApp encryptor_applications.Application,
	adapter accounts.Adapter,
	dbPtr *sql.DB,
) accounts.Repository {
	return createAccountRepository(
		encryptorApp,
		adapter,
		dbPtr,
	)
}

// NewAccountService creates a new account service
func NewAccountService(
	encryptorApp encryptor_applications.Application,
	repository accounts.Repository,
	adapter accounts.Adapter,
	bitrate int,
	dbPtr *sql.DB,
) accounts.Service {
	builder := accounts.NewBuilder()
	encryptorBuilder := encryptors.NewBuilder()
	signerFactory := signers.NewFactory()
	return createAccountService(
		encryptorApp,
		builder,
		repository,
		adapter,
		encryptorBuilder,
		signerFactory,
		bitrate,
		dbPtr,
	)
}
