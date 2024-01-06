package threads

import "steve.care/network/domain/hash"

type threads struct {
	hash hash.Hash
	list []Thread
}

func createThreads(
	hash hash.Hash,
	list []Thread,
) Threads {
	out := threads{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *threads) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *threads) List() []Thread {
	return obj.list
}
