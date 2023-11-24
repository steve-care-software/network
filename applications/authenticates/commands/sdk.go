package commands

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands/links"
)

// Application represents the command authenticated application
type Application interface {
	Exists(hash hash.Hash) (bool, error)
	Execute(hash hash.Hash, input []byte) (receipts.Receipt, error)
	Links(executed []hash.Hash) (links.Link, error)
	Clear() error
}
