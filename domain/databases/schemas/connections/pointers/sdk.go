package pointers

// Builder represents the pointer builder
type Builder interface {
	Create() Builder
	WithEntity(entity string) Builder
	WithField(field string) Builder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	Entity() string
	Field() string
}
