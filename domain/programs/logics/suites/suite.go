package suites

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
	"steve.care/network/domain/programs/logics/suites/expectations"
)

type suite struct {
	hash        hash.Hash
	origin      links.Origin
	input       []byte
	expectation expectations.Expectation
}

func createSuite(
	hash hash.Hash,
	origin links.Origin,
	input []byte,
	expectation expectations.Expectation,
) Suite {
	out := suite{
		hash:        hash,
		origin:      origin,
		input:       input,
		expectation: expectation,
	}

	return &out
}

// Hash returns the hash
func (obj *suite) Hash() hash.Hash {
	return obj.hash
}

// Origin returns the origin
func (obj *suite) Origin() links.Origin {
	return obj.origin
}

// Input returns the input
func (obj *suite) Input() []byte {
	return obj.input
}

// Expectation returns the expectation
func (obj *suite) Expectation() expectations.Expectation {
	return obj.expectation
}
