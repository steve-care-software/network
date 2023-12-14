package receipts

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a receipt builder
type Builder interface {
	Create() Builder
	WithCommands(commands commands.Commands) Builder
	WithSignature(signature signers.Signature) Builder
	Now() (Receipt, error)
}

// Receipt represents a receipt
type Receipt interface {
	Hash() hash.Hash
	Commands() commands.Commands
	Signature() signers.Signature
}
