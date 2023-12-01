package accounts

import (
	"steve.care/network/domain/accounts"
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases"
	"steve.care/network/domain/encryptors"
)

// NewBuilder creates a new bulder
func NewBuilder(
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
) Builder {
	accountBuilder := accounts.NewBuilder()
	signerFactory := signers.NewFactory()
	encryptorBuilder := account_encryptors.NewBuilder()
	repositoryBuilder := accounts.NewRepositoryBuilder(
		encryptor,
		adapter,
	)

	serviceBuilder := accounts.NewServiceBuilder(
		encryptor,
		adapter,
	)

	return createBuilder(
		accountBuilder,
		signerFactory,
		encryptorBuilder,
		repositoryBuilder,
		serviceBuilder,
	)
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithDatabase(database databases.Database) Builder
	WithBitrate(bitrate int) Builder
	Now() (Application, error)
}

// Application represents the authenticated account application
type Application interface {
	List() ([]string, error)
	Exists(username string) (bool, error)
	Insert(credentials credentials.Credentials) error
	Retrieve(credentials credentials.Credentials) (accounts.Account, error)
	Update(credentials credentials.Credentials, criteria accounts.UpdateCriteria) error
	Delete(credentials credentials.Credentials) error
}
