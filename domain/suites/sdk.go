package suites

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers"
	"steve.care/network/domain/receipts/commands/links"
	"steve.care/network/domain/suites/expectations"
)

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
