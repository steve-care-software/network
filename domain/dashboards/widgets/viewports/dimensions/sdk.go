package dimensions

import (
	"steve.care/network/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the dimension builder
type Builder interface {
	Create() Builder
	WithWidth(width float32) Builder
	WithHeight(height float32) Builder
	Now() (Dimension, error)
}

// Dimension represents a dimension
type Dimension interface {
	Hash() hash.Hash
	Width() float32
	Height() float32
}
