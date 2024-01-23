package expectations

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
	"steve.care/network/domain/programs/logics/suites/expectations/outputs"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an expectation builder
type Builder interface {
	Create() Builder
	WithSuccess(success outputs.Output) Builder
	WithMistake(mistake links.Condition) Builder
	Now() (Expectation, error)
}

// Expectation represents an expectation
type Expectation interface {
	Hash() hash.Hash
	IsSuccess() bool
	Success() outputs.Output
	IsMistake() bool
	Mistake() links.Condition
}
