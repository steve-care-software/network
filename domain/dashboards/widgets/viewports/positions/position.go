package positions

import "steve.care/network/domain/hash"

type position struct {
	hash       hash.Hash
	horizontal float32
	vertical   float32
}

func createPosition(
	hash hash.Hash,
	horizontal float32,
	vertical float32,
) Position {
	out := position{
		hash:       hash,
		horizontal: horizontal,
		vertical:   vertical,
	}

	return &out
}

// Hash returns the hash
func (obj *position) Hash() hash.Hash {
	return obj.hash
}

// Horizontal returns the horizontal
func (obj *position) Horizontal() float32 {
	return obj.horizontal
}

// Vertical returns the vertical
func (obj *position) Vertical() float32 {
	return obj.vertical
}
