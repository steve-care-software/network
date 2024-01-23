package schemas

import "steve.care/network/domain/schemas/resources"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a schema builder
type Builder interface {
	Create() Builder
	WithResources(resources resources.Resources) Builder
	WithPrevious(previous Schema) Builder
	Now() (Schema, error)
}

// Schema represents a schema
type Schema interface {
	Version() uint
	Resources() resources.Resources
	HasPrevious() bool
	Previous() Schema
}
