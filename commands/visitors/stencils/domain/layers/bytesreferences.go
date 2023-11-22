package layers

import (
	"steve.care/network/libraries/hash"
)

type bytesReferences struct {
	hash hash.Hash
	list []BytesReference
}

func createBytesReferences(
	hash hash.Hash,
	list []BytesReference,
) BytesReferences {
	out := bytesReferences{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *bytesReferences) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *bytesReferences) List() []BytesReference {
	return obj.list
}
