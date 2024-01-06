package logics

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/suites"
	"steve.care/network/domain/programs/logics/threads"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the logic builder
type Builder interface {
	Create() Builder
	WithEntry(entry layers.Layer) Builder
	WithLibrary(library libraries.Library) Builder
	WithSuites(suites suites.Suites) Builder
	WithThreads(threads threads.Threads) Builder
	Now() (Logic, error)
}

// Logic represents the program's logic
type Logic interface {
	Hash() hash.Hash
	Entry() layers.Layer
	Library() libraries.Library
	HasSuites() bool
	Suites() suites.Suites
	HasThreads() bool
	Threads() threads.Threads
}
