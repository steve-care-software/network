package programs

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/logics"
)

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithLogic(logic logics.Logic) Builder
	WithHead(head blocks.Block) Builder
	WithParent(parent Program) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Name() string
	Description() string
	HasHead() bool
	Head() blocks.Block
	HasLogic() bool
	Logic() logics.Logic
	HasParent() bool
	Parent() Program
}
