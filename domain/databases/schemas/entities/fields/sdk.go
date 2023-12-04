package fields

import "steve.care/network/domain/databases/schemas/entities/fields/kinds"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewFieldBuilder creates a new field instance
func NewFieldBuilder() FieldBuilder {
	return createFieldBuilder()
}

// Builder represents the fields builder
type Builder interface {
	Create() Builder
	WithList(list []Field) Builder
	Now() (Fields, error)
}

// Fields represents fields
type Fields interface {
	List() []Field
}

// FieldBuilder represents the field builder
type FieldBuilder interface {
	Create() FieldBuilder
	WithName(name string) FieldBuilder
	WithKind(kind kinds.Kind) FieldBuilder
	IsUnique() FieldBuilder
	Now() (Field, error)
}

// Field represents a field
type Field interface {
	Name() string
	Kind() kinds.Kind
	IsUnique() bool
}
