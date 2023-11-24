package receipts

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands"
)

// Builder represents a receipt builder
type Builder interface {
	Create() Builder
	WithCommand(command commands.Commands) Builder
	WithSignature(signature signers.Signature) Builder
	Now() (Receipt, error)
}

// Receipt represents a receipt
type Receipt interface {
	Hash() hash.Hash
	Command() commands.Commands
	Signature() signers.Signature
}
