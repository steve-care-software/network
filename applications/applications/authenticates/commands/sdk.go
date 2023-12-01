package commands

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands/results"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the command authenticated application
type Application interface {
	Begin() (*uint, error)
	Execute(context uint, hash hash.Hash, input []byte) (results.Result, error)
	End(context uint) (receipts.Receipt, error)
}
