package methods

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a methods builder
type Builder interface {
	Create() Builder
	WithInitialize(initialize string) Builder
	WithTrigger(trigger string) Builder
	WithElement(element string) Builder
	Now() (Methods, error)
}

// Methods represents a methods
type Methods interface {
	Initialize() string
	Trigger() string
	HasElement() bool
	Element() string
}
