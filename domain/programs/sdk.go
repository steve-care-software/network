package programs

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/logics"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithSpace(space []string) Builder
	WithDescription(description string) Builder
	WithLogic(logic logics.Logic) Builder
	WithHead(head blocks.Block) Builder
	WithParent(parent Program) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Space() []string
	Description() string
	HasHead() bool
	Head() blocks.Block
	HasLogic() bool
	Logic() logics.Logic
	HasParent() bool
	Parent() Program
}
