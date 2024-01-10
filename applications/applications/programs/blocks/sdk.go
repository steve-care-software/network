package blocks

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/blocks/transactions"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	WithDifficulty(difficulty blocks.Difficulty) Builder
	Now() (Application, error)
}

// Application represents the block application
type Application interface {
	Exists(hash hash.Hash) bool
	Retrieve(hash hash.Hash) (blocks.Block, error)
	RetrieveHeadByProgram(program hash.Hash) (blocks.Block, error)
	Transact(trx transactions.Transactions) error
	Queue() (transactions.Transactions, error)
	Mine(program hash.Hash) error
	Rewind(head hash.Hash) error
	Insert(blocks blocks.Block) error
}
