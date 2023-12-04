package suites

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/suites"
	"steve.care/network/domain/suites/expectations"
)

// Suite represents a suite resource
type Suite interface {
	Hash() hash.Hash
	IsSuite() bool
	Suite() suites.Suite
	IsExpectation() bool
	Expectation() expectations.Expectation
}
