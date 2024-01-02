package expectations

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers"
	"steve.care/network/domain/receipts/commands/layers/links"
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
	WithOutput(output layers.Layer) Builder
	WithCondition(condition links.Condition) Builder
	Now() (Expectation, error)
}

// Expectation represents an expectation
type Expectation interface {
	Hash() hash.Hash
	IsOutput() bool
	Output() layers.Layer
	IsCondition() bool
	Condition() links.Condition
}
