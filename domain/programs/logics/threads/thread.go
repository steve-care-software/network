package threads

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

type thread struct {
	hash  hash.Hash
	input []byte
	entry layers.Layer
}

func createThread(
	hash hash.Hash,
	input []byte,
	entry layers.Layer,
) Thread {
	out := thread{
		hash:  hash,
		input: input,
		entry: entry,
	}

	return &out
}

// Hash returns the hash
func (obj *thread) Hash() hash.Hash {
	return obj.hash
}

// Input returns the input
func (obj *thread) Input() []byte {
	return obj.input
}

// Entry returns the entry
func (obj *thread) Entry() layers.Layer {
	return obj.entry
}
