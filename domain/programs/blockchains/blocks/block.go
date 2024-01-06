package blocks

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains/blocks/executions"
)

type block struct {
	hash       hash.Hash
	executions executions.Executions
	parent     hash.Hash
}

func createBlock(
	hash hash.Hash,
	executions executions.Executions,
) Block {
	return createBlockInternally(hash, executions, nil)
}

func createBlockWithParent(
	hash hash.Hash,
	executions executions.Executions,
	parent hash.Hash,
) Block {
	return createBlockInternally(hash, executions, parent)
}

func createBlockInternally(
	hash hash.Hash,
	executions executions.Executions,
	parent hash.Hash,
) Block {
	out := block{
		hash:       hash,
		executions: executions,
		parent:     parent,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.hash
}

// Executions returns the executions
func (obj *block) Executions() executions.Executions {
	return obj.executions
}

// HasParent returns true if there is parent, false otherwise
func (obj *block) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *block) Parent() hash.Hash {
	return obj.parent
}
