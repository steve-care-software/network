package blockchains

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains/blocks"
	"steve.care/network/domain/programs/blockchains/roots"
)

// Blockchain represents a blockchain
type Blockchain interface {
	Hash() hash.Hash
	Root() roots.Root
	Head() blocks.Block
}
