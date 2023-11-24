package results

import "steve.care/network/domain/hash"

type failure struct {
	hash            hash.Hash
	code            uint
	isRaisedInLayer bool
	pIndex          *uint
}

func createFailure(
	hash hash.Hash,
	code uint,
	isRaisedInLayer bool,
) Failure {
	return createFailureInternally(hash, code, isRaisedInLayer, nil)
}

func createFailureWithIndex(
	hash hash.Hash,
	code uint,
	isRaisedInLayer bool,
	pIndex *uint,
) Failure {
	return createFailureInternally(hash, code, isRaisedInLayer, pIndex)
}

func createFailureInternally(
	hash hash.Hash,
	code uint,
	isRaisedInLayer bool,
	pIndex *uint,
) Failure {
	out := failure{
		hash:            hash,
		code:            code,
		isRaisedInLayer: isRaisedInLayer,
		pIndex:          pIndex,
	}

	return &out
}

// Hash returns the hash
func (obj *failure) Hash() hash.Hash {
	return obj.hash
}

// Code returns the code
func (obj *failure) Code() uint {
	return obj.code
}

// IsRaisedInLayer returns true if raisedInLayer, false otherwise
func (obj *failure) IsRaisedInLayer() bool {
	return obj.isRaisedInLayer
}

// HasIndex returns true if there is an index, false otherwise
func (obj *failure) HasIndex() bool {
	return obj.pIndex != nil
}

// Index returns the index, if any
func (obj *failure) Index() *uint {
	return obj.pIndex
}
