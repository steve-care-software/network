package threads

import (
	"time"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

type thread struct {
	hash       hash.Hash
	input      []byte
	entry      layers.Layer
	waitPeriod time.Duration
}

func createThread(
	hash hash.Hash,
	input []byte,
	entry layers.Layer,
	waitPeriod time.Duration,
) Thread {
	out := thread{
		hash:       hash,
		input:      input,
		entry:      entry,
		waitPeriod: waitPeriod,
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

// WaitPeriod returns the waitPeriod
func (obj *thread) WaitPeriod() time.Duration {
	return obj.waitPeriod
}
