package dimensions

import "steve.care/network/domain/hash"

type dimension struct {
	hash   hash.Hash
	width  float32
	height float32
}

func createDimension(
	hash hash.Hash,
	width float32,
	height float32,
) Dimension {
	out := dimension{
		hash:   hash,
		width:  width,
		height: height,
	}

	return &out
}

// Hash returns the hash
func (obj *dimension) Hash() hash.Hash {
	return obj.hash
}

// Width returns the width
func (obj *dimension) Width() float32 {
	return obj.width
}

// Height returns the height
func (obj *dimension) Height() float32 {
	return obj.height
}
