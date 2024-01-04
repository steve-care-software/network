package blocks

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains/blocks/executions"
)

// Block representa a block
type Block interface {
	Hash() hash.Hash
	Executions() executions.Executions
	HasParent() bool
	Parent() hash.Hash
}
