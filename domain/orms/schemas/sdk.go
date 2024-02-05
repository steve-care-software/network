package schemas

import "steve.care/network/domain/orms/schemas/roots"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Factory represents a schema factory
type Factory interface {
	Create() (Schema, error)
}

// Builder represents a schema builder
type Builder interface {
	Create() Builder
	WithRoots(roots roots.Roots) Builder
	WithPrevious(previous Schema) Builder
	Now() (Schema, error)
}

// Schema represents a schema
type Schema interface {
	Version() uint
	Roots() roots.Roots
	HasPrevious() bool
	Previous() Schema
}
