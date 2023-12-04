package resources

import (
	"steve.care/network/domain/databases/criterias/conditions"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a resource builder
type Builder interface {
	Create() Builder
	WithContainer(container string) Builder
	WithCondition(condition conditions.Condition) Builder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Container() string
	Condition() conditions.Condition
}
