package outputs

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an output builder
type Builder interface {
	Create() Builder
	WithKind(kind layers.Kind) Builder
	WithValue(value []byte) Builder
	Now() (Output, error)
}

// Output represents the output
type Output interface {
	Hash() hash.Hash
	Kind() layers.Kind
	Value() []byte
}
