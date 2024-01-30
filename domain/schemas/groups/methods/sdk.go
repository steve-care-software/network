package methods

import (
	resource_methods "steve.care/network/domain/schemas/groups/resources/methods"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the methods builder
type Builder interface {
	Create() Builder
	WithBuilder(builder resource_methods.Methods) Builder
	Now() (Methods, error)
}

// Methods represents methods
type Methods interface {
	Builder() resource_methods.Methods
}
