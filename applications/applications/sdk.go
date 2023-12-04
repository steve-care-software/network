package applications

import (
	accounts_application "steve.care/network/applications/applications/accounts"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/databases/schemas"
	"steve.care/network/domain/databases/transactions"
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

	return createBuilder(
		accAppBuilder,
	)

}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithQuery(query queries.Query) Builder
	WithTransaction(trx transactions.Transaction) Builder
	WithBitrate(bitrate int) Builder
	Now() (Application, error)
}

// Application represents a stencil application
type Application interface {
	Schema() (schemas.Schema, error)
	Accounts() (accounts_application.Application, error)
}
