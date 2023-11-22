package layers

import (
	"steve.care/network/libraries/hash"
)

type bytesReference struct {
	hash     hash.Hash
	variable string
	bytes    []byte
}

func createBytesReferenceWithVariable(
	hash hash.Hash,
	variable string,
) BytesReference {
	return createBytesReferenceInternally(hash, variable, nil)
}

func createBytesReferenceWithBytes(
	hash hash.Hash,
	bytes []byte,
) BytesReference {
	return createBytesReferenceInternally(hash, "", bytes)
}

func createBytesReferenceInternally(
	hash hash.Hash,
	variable string,
	bytes []byte,
) BytesReference {
	out := bytesReference{
		hash:     hash,
		variable: variable,
		bytes:    bytes,
	}

	return &out
}

// Hash returns the hash
func (obj *bytesReference) Hash() hash.Hash {
	return obj.hash
}

// IsVariable returns true if there is a variable, false otherwise
func (obj *bytesReference) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *bytesReference) Variable() string {
	return obj.variable
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *bytesReference) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *bytesReference) Bytes() []byte {
	return obj.bytes
}
