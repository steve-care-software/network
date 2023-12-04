package criterias

import (
	"steve.care/network/domain/databases/criterias/conditions"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
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
	Entity() string
	Condition() conditions.Condition
}
