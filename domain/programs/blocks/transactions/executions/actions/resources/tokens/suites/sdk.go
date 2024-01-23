package suites

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/suites"
	"steve.care/network/domain/programs/logics/suites/expectations"
	"steve.care/network/domain/programs/logics/suites/expectations/outputs"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the suite builder
type Builder interface {
	Create() Builder
	WithSuite(suite suites.Suite) Builder
	WithExpectation(expectation expectations.Expectation) Builder
	WithOutput(output outputs.Output) Builder
	Now() (Suite, error)
}

// Suite represents a suite resource
type Suite interface {
	Hash() hash.Hash
	IsSuite() bool
	Suite() suites.Suite
	IsExpectation() bool
	Expectation() expectations.Expectation
	IsOutput() bool
	Output() outputs.Output
}
