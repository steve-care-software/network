package criterias

import (
	"steve.care/network/domain/criterias/conditions"
	"steve.care/network/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a criteria builder
type Builder interface {
	Create() Builder
	WithEntity(entity string) Builder
	WithCondition(condition conditions.Condition) Builder
	Now() (Criteria, error)
}

// Criteria represents a criteria
type Criteria interface {
	Hash() hash.Hash
	Entity() string
	Condition() conditions.Condition
}
