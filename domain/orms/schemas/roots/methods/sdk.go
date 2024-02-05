package methods

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the methods builder
type Builder interface {
	Create() Builder
	WithInitialize(initialize string) Builder
	WithTrigger(trigger string) Builder
	Now() (Methods, error)
}

// Methods represents methods
type Methods interface {
	Initialize() string
	Trigger() string
}
