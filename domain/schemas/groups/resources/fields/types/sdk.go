package types

import "steve.care/network/domain/schemas/groups/resources/fields/types/dependencies"

const (
	// KindInteger represents the integer kind
	KindInteger (uint8) = iota

	// KindFloat represents the float kind
	KindFloat

	// KindString represents the string kind
	KindString

	// KindBytes represents the bytes kind
	KindBytes
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a type builder
type Builder interface {
	Create() Builder
	WithKind(kind uint8) Builder
	WithDependency(dependency dependencies.Dependency) Builder
	Now() (Type, error)
}

// Type represents a type
type Type interface {
	IsKind() bool
	Kind() *uint8
	IsDependency() bool
	Dependency() dependencies.Dependency
}
