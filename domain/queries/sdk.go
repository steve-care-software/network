package queries

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/queries/conditions"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a query builder
type Builder interface {
	Create() Builder
	WithEntity(entity string) Builder
	WithCondition(condition conditions.Condition) Builder
	WithFields(fields []string) Builder
	Now() (Query, error)
}

// Query represents a query
type Query interface {
	Hash() hash.Hash
	Entity() string
	Condition() conditions.Condition
	HasFields() bool
	Fields() []string
}
