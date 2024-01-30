package methods

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a methods builder
type Builder interface {
	Create() Builder
	WithRetriever(retriever []string) Builder
	WithBuilder(builder string) Builder
	Now() (Methods, error)
}

// Methods represents field methods
type Methods interface {
	Retriever() []string
	Builder() string
}
