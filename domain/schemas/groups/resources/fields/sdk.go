package fields

import (
	"steve.care/network/domain/schemas/groups/resources/fields/methods"
	"steve.care/network/domain/schemas/groups/resources/fields/types"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewFieldBuilder creates a new field builder
func NewFieldBuilder() FieldBuilder {
	return createFieldBuilder()
}

// Builder represents a fields builder
type Builder interface {
	Create() Builder
	WithList(list []Field) Builder
	Now() (Fields, error)
}

// Fields represents fields
type Fields interface {
	List() []Field
}

// FieldBuilder represents a field builder
type FieldBuilder interface {
	Create() FieldBuilder
	WithName(name string) FieldBuilder
	WithMethods(methods methods.Methods) FieldBuilder
	WithType(typ types.Type) FieldBuilder
	CanBeNil() FieldBuilder
	Now() (Field, error)
}

// Field represents a field
type Field interface {
	Name() string
	Methods() methods.Methods
	Type() types.Type
	CanBeNil() bool
}
