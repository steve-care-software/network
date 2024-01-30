package methods

import (
	field_methods "steve.care/network/domain/schemas/groups/resources/fields/methods"
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
	WithField(field field_methods.Methods) Builder
	Now() (Methods, error)
}

// Methods represents a methods
type Methods interface {
	Initialize() string
	Trigger() string
	Field() field_methods.Methods
}
