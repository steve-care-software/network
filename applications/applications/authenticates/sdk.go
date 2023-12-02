package authenticates

import (
	commands_application "steve.care/network/applications/applications/authenticates/commands"
	layers_application "steve.care/network/applications/applications/authenticates/layers"
	links_application "steve.care/network/applications/applications/authenticates/links"
	receipts_application "steve.care/network/applications/applications/authenticates/receipts"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/databases/transactions"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	WithQuery(query queries.Query) Builder
	WithTransaction(trx transactions.Transaction) Builder
	Now() (Application, error)
}

// Application represents an authenticated application
type Application interface {
	Receipts() receipts_application.Application
	Layers() layers_application.Application
	Links() links_application.Application
	Commands() commands_application.Application
}
