package viewports

import (
	"steve.care/network/domain/hash"
)

type viewport struct {
	hash   hash.Hash
	row    uint
	height uint
}

func createViewport(
	hash hash.Hash,
	row uint,
	height uint,
) Viewport {
	out := viewport{
		hash:   hash,
		row:    row,
		height: height,
	}

	return &out
}

// Hash returns the hash
func (obj *viewport) Hash() hash.Hash {
	return obj.hash
}

// Row returns the row
func (obj *viewport) Row() uint {
	return obj.row
}

// Height returns the height
func (obj *viewport) Height() uint {
	return obj.height
}
