package layers

import "steve.care/network/libraries/hash"

type assignment struct {
	hash       hash.Hash
	name       string
	assignable Assignable
}

func createAssignment(
	hash hash.Hash,
	name string,
	assignable Assignable,
) Assignment {
	out := assignment{
		hash:       hash,
		name:       name,
		assignable: assignable,
	}

	return &out
}

// Hash returns the hash
func (obj *assignment) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *assignment) Name() string {
	return obj.name
}

// Assignable returns the assignable
func (obj *assignment) Assignable() Assignable {
	return obj.assignable
}
