package fields

const (
	// KindNil represents the nil kind
	KindNil (uint8) = iota

	// KindInteger represents the integer kind
	KindInteger

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
	WithKind(kind uint8) FieldBuilder
	CanBeNil() FieldBuilder
	Now() (Field, error)
}

// Field represents a field
type Field interface {
	Name() string
	Kind() uint8
	CanBeNil() bool
}
