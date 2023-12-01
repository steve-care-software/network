package applications

import (
	accounts_application "steve.care/network/applications/applications/accounts"
	"steve.care/network/applications/applications/authenticates"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases"
	"steve.care/network/domain/encryptors"
)

// NewBuilder creates a new appication builder
func NewBuilder(
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
) Builder {
	accAppBuilder := accounts_application.NewBuilder(
		encryptor,
		adapter,
	)

	var authAppBuilder authenticates.Builder
	return createBuilder(
		accAppBuilder,
		authAppBuilder,
	)

}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithQuery(query databases.Query) Builder
	WithTransaction(trx databases.Transaction) Builder
	WithBitrate(bitrate int) Builder
	Now() (Application, error)
}

// Application represents a stencil application
type Application interface {
	Accounts() (accounts_application.Application, error)
	Authenticate(credentials credentials.Credentials) (authenticates.Application, error)
}
