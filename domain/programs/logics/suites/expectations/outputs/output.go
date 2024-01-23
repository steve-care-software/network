package outputs

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

type output struct {
	hash  hash.Hash
	kind  layers.Kind
	value []byte
}

func createOutput(
	hash hash.Hash,
	kind layers.Kind,
	value []byte,
) Output {
	out := output{
		hash:  hash,
		kind:  kind,
		value: value,
	}

	return &out
}

// Hash returns the hash
func (obj *output) Hash() hash.Hash {
	return obj.hash
}

// Kind returns the kind
func (obj *output) Kind() layers.Kind {
	return obj.kind
}

// Value returns the value
func (obj *output) Value() []byte {
	return obj.value
}
