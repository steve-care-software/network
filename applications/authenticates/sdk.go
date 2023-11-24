package authenticates

import (
	accounts_application "steve.care/network/applications/authenticates/accounts"
	commands_appication "steve.care/network/applications/authenticates/commands"
	receipts_application "steve.care/network/applications/authenticates/receipts"
	"steve.care/network/domain/accounts"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithAccount(account accounts.Account) Builder
	NOw() (Application, error)
}

// Application represents an authenticated application
type Application interface {
	Account() accounts_application.Application
	Receipt() receipts_application.Application
	Command() commands_appication.Application
}
