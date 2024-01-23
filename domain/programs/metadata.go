package programs

import "steve.care/network/domain/hash"

type metadata struct {
	hash   hash.Hash
	name   string
	parent Program
}

func createMetaData(
	hash hash.Hash,
	name string,
	parent Program,
) MetaData {
	out := metadata{
		hash:   hash,
		name:   name,
		parent: parent,
	}

	return &out
}

// Hash returns the hash
func (obj *metadata) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *metadata) Name() string {
	return obj.name
}

// Parent returns the parent
func (obj *metadata) Parent() Program {
	return obj.parent
}
