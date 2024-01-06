package positions

import (
	"steve.care/network/domain/hash"
)

// Position represents a position
type Position interface {
	Hash() hash.Hash
	Horizontal() float32
	Vertical() float32
}
