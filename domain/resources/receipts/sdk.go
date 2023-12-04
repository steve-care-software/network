package receipts

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands"
)

// Receipt represents a receipt resource
type Receipt interface {
	Hash() hash.Hash
	IsReceipt() bool
	Receipt() receipts.Receipt
	IsCommand() bool
	Command() commands.Command
	IsLink() bool
	Link() commands.Link
}
