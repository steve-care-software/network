package commands

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands/results"
)

// Application represents the command authenticated application
type Application interface {
	Begin() (*uint, error)
	Execute(context uint, credentials credentials.Credentials, hash hash.Hash, input []byte) (results.Result, error)
	End(context uint) (receipts.Receipt, error)
}
