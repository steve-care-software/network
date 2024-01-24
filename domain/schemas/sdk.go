package schemas

import "steve.care/network/domain/schemas/groups"

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
	WithGroups(groups groups.Groups) Builder
	WithPrevious(previous Schema) Builder
	Now() (Schema, error)
}

// Schema represents a schema
type Schema interface {
	Version() uint
	Groups() groups.Groups
	HasPrevious() bool
	Previous() Schema
}
