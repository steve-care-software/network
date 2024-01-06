package blocks

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains/blocks/executions"
)

// Builder represents the block builder
type Builder interface {
	Create() Builder
	WithExecutions(executions executions.Executions) Builder
	WithParent(parent hash.Hash) Builder
	Now() (Block, error)
}

// Block representa a block
type Block interface {
	Hash() hash.Hash
	Executions() executions.Executions
	HasParent() bool
	Parent() hash.Hash
}
