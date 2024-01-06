package blockchains

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains/blocks"
	"steve.care/network/domain/programs/blockchains/permissions"
)

// Builder represents the blockchain builder
type Builder interface {
	Create() Builder
	WithHead(head blocks.Block) Builder
	WithPermission(permission permissions.Permission) Builder
	Now() (Blockchain, error)
}

// Blockchain represents a blockchain
type Blockchain interface {
	Hash() hash.Hash
	Head() blocks.Block
	HasPermission() bool
	Permission() permissions.Permission
}
