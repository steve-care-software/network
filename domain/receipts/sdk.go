package receipts

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands"
)

// NewReceiptBuilder creates a new builder instance
func NewReceiptBuilder() ReceiptBuilder {
	hashAdapter := hash.NewAdapter()
	return createReceiptBuilder(
		hashAdapter,
	)
}

// ReceiptBuilder represents a receipt builder
type ReceiptBuilder interface {
	Create() ReceiptBuilder
	WithCommands(commands commands.Commands) ReceiptBuilder
	WithSignature(signature signers.Signature) ReceiptBuilder
	Now() (Receipt, error)
}

// Receipt represents a receipt
type Receipt interface {
	Hash() hash.Hash
	Commands() commands.Commands
	Signature() signers.Signature
}
