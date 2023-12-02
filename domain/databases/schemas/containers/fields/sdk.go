package fields

import "steve.care/network/domain/databases/schemas/containers/fields/kinds"

// Fields represents fields
type Fields interface {
	List() []Field
}

// Field represents a field
type Field interface {
	Name() string
	Kind() kinds.Kind
	IsUnique() bool
}
