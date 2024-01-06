package logics

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

// Builder represents the logic builder
type Builder interface {
	Create() Builder
	WithEntry(entry layers.Layer) Builder
	WithLibrary(library libraries.Library) Builder
	Now() (Logic, error)
}

// Logic represents the program's logic
type Logic interface {
	Hash() hash.Hash
	Entry() layers.Layer
	HasLibrary() bool
	Library() libraries.Library
}
