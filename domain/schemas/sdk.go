package schemas

import "steve.care/network/domain/schemas/roots"

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
	WithRoot(root roots.Root) Builder
	WithPrevious(previous Schema) Builder
	Now() (Schema, error)
}

// Schema represents a schema
type Schema interface {
	Version() uint
	Root() roots.Root
	HasPrevious() bool
	Previous() Schema
}
