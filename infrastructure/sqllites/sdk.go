package sqllites

import (
	"database/sql"

	"steve.care/network/applications"
	"steve.care/network/domain/accounts"
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/encryptors"
)

const notActiveErrorMsg = "the application NEVER began a transactional state, therefore that method cannot be executed"
const currentActiveErrorMsg = "the application ALREADY began a transactional state, therefore that method cannot be executed"

// NewApplication creates a new application
func NewApplication(
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	bitrate int,
	basePath string,
) applications.Application {
	return createApplication(
		encryptor,
		adapter,
		bitrate,
		basePath,
	)
}

// NewAccountRepository creates a new account repository
func NewAccountRepository(
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	dbPtr *sql.DB,
) accounts.Repository {
	return createAccountRepository(
		encryptor,
		adapter,
		dbPtr,
	)
}

// NewAccountService creates a new account service
func NewAccountService(
	repository accounts.Repository,
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	bitrate int,
	txPtr *sql.Tx,
) accounts.Service {
	builder := accounts.NewBuilder()
	encryptorBuilder := account_encryptors.NewBuilder()
	signerFactory := signers.NewFactory()
	return createAccountService(
		encryptor,
		builder,
		repository,
		adapter,
		encryptorBuilder,
		signerFactory,
		bitrate,
		txPtr,
	)
}
