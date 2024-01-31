package dependencies

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a dependency builder
type Builder interface {
	Create() Builder
	WithRetriever(retriever string) Builder
	WithGroups(groups []string) Builder
	WithResource(resource string) Builder
	WithKind(kind uint8) Builder
	Now() (Dependency, error)
}

// Dependency represents a dependency
type Dependency interface {
	Retriever() string
	Groups() []string
	Resource() string
	Kind() uint8
}
