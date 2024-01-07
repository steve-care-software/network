package viewports

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

// Builder represents the viewport builder
type Builder interface {
	Create() Builder
	WithRow(row uint) Builder
	WithHeight(height uint) Builder
	Now() (Viewport, error)
}

// Viewport represents a viewport
type Viewport interface {
	Hash() hash.Hash
	Row() uint
	Height() uint
}
