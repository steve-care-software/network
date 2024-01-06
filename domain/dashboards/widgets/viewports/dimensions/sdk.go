package dimensions

import (
	"steve.care/network/domain/hash"
)

// Dimension represents a dimension
type Dimension interface {
	Hash() hash.Hash
	Width() float32
	Height() float32
}
