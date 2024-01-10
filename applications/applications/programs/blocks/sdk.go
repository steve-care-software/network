package blocks

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/blocks/executions"
)

// Application represents the block application
type Application interface {
	Retrieve(hash hash.Hash) (blocks.Block, error)
	RetrieveHeadByProgram(program hash.Hash) (blocks.Block, error)
	Queue() (executions.Executions, error)
	Mine(program hash.Hash) (blocks.Block, error)
	Rewind(head hash.Hash) (blocks.Block, error)
	Insert(block blocks.Block) error
	Save(block blocks.Block) error
}
