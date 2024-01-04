package receipts

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands"
	"steve.care/network/domain/receipts/commands/results"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a receipt builder
type Builder interface {
	Create() Builder
	WithReceipt(receipt receipts.Receipt) Builder
	WithCommand(command commands.Command) Builder
	WithResult(result results.Result) Builder
	WithSuccess(success results.Success) Builder
	WithFailure(failure results.Failure) Builder
	WithLink(link commands.Link) Builder
	Now() (Receipt, error)
}

// Receipt represents a receipt resource
type Receipt interface {
	Hash() hash.Hash
	IsReceipt() bool
	Receipt() receipts.Receipt
	IsCommand() bool
	Command() commands.Command
	IsResult() bool
	Result() results.Result
	IsSuccess() bool
	Success() results.Success
	IsFailure() bool
	Failure() results.Failure
	IsLink() bool
	Link() commands.Link
}
