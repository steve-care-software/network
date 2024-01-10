package blocks

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the block builder
type Builder interface {
	Create() Builder
	WithTransactions(transactions transactions.Transactions) Builder
	WithParent(parent hash.Hash) Builder
	Now() (Block, error)
}

// Block representa a block
type Block interface {
	Hash() hash.Hash
	Transactions() transactions.Transactions
	HasParent() bool
	Parent() hash.Hash
}
