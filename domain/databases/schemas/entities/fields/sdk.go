package fields

import "steve.care/network/domain/databases/schemas/entities/fields/kinds"

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
