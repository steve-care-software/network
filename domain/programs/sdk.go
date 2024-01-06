package programs

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains"
	"steve.care/network/domain/programs/logics"
)

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithLogic(logic logics.Logic) Builder
	WithBlockchain(blockchain blockchains.Blockchain) Builder
	WithParent(parent Program) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Logic() logics.Logic
	Blockchain() blockchains.Blockchain
	HasParent() bool
	Parent() Program
}
