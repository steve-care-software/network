package suites

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
	"steve.care/network/domain/programs/logics/suites/expectations"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Suites represents suites
type Suites interface {
	Hash() hash.Hash
	List() []Suite
}

// Builder represents a suite builder
type Builder interface {
	Create() Builder
	WithOrigin(origin links.Origin) Builder
	WithInput(input layers.Layer) Builder
	WithExpectation(expectation expectations.Expectation) Builder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Hash() hash.Hash
	Origin() links.Origin
	Input() layers.Layer
	Expectation() expectations.Expectation
}
