package positions

import (
	"steve.care/network/domain/hash"
)

// Builder represents the position builder
type Builder interface {
	Create() Builder
	WithHorizontal(horizontal float32) Builder
	WithVertical(vertical float32) Builder
	Now() (Position, error)
}

// Position represents a position
type Position interface {
	Hash() hash.Hash
	Horizontal() float32
	Vertical() float32
}
