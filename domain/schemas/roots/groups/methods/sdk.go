package methods

import (
	root_methods "steve.care/network/domain/schemas/roots/methods"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a methods builder
type Builder interface {
	Create() Builder
	WithInitialize(initialize string) Builder
	WithTrigger(trigger string) Builder
	WithElement(element string) Builder
	Now() (Methods, error)
}

// Methods represents a methods
type Methods interface {
	root_methods.Methods
	Element() string
}
