package dependencies

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a dependency builder
type Builder interface {
	Create() Builder
	WithGroups(groups []string) Builder
	WithResource(resource string) Builder
	Now() (Dependency, error)
}

// Dependency represents a dependency
type Dependency interface {
	Groups() []string
	Resource() string
}
