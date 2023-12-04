package kinds

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the kind builder
type Builder interface {
	Create() Builder
	IsNil() Builder
	IsInteger() Builder
	IsReal() Builder
	IsText() Builder
	IsBlob() Builder
	Now() (Kind, error)
}

// Kind represents the kind of field
type Kind interface {
	IsNil() bool
	IsInteger() bool
	IsReal() bool
	IsText() bool
	IsBlob() bool
}
