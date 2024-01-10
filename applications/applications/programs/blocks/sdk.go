package blocks

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/blocks/transactions"
)

// Application represents the block application
type Application interface {
	Retrieve(hash hash.Hash) (blocks.Block, error)
	RetrieveHeadByProgram(program hash.Hash) (blocks.Block, error)
	Transact(trx transactions.Transactions) error
	Queue() (transactions.Transactions, error)
	Mine(program hash.Hash) (blocks.Block, error)
	Rewind(head hash.Hash) (blocks.Block, error)
	Insert(block blocks.Block) error
	Save(block blocks.Block) error
}
