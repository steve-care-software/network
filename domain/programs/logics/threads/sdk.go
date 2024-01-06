package threads

import (
	"time"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

// Builder represents a threads builder
type Builder interface {
	Create() Builder
	WithList(list []Thread) Builder
	Now() (Threads, error)
}

// Threads represents threads
type Threads interface {
	Hash() hash.Hash
	List() []Thread
}

// ThreadBuilder represents a thread builder
type ThreadBuilder interface {
	Create() ThreadBuilder
	WithInput(input []byte) ThreadBuilder
	WithEntry(entry layers.Layer) ThreadBuilder
	WithWaitPeriod(waitPeriod time.Duration) ThreadBuilder
	Now() (Thread, error)
}

// Thread represents a thread
type Thread interface {
	Hash() hash.Hash
	Input() []byte
	Entry() layers.Layer
	WaitPeriod() time.Duration
}
