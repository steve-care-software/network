package programs

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains"
	"steve.care/network/domain/programs/logics"
)

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Logic() logics.Logic
	Blockchain() blockchains.Blockchain
	HasParent() bool
	Parent() Program
}
