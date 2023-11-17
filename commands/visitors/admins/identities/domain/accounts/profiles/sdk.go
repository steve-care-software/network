package profiles

// Builder represents a profile builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	Now() (Profile, error)
}

// Profile represents a profile
type Profile interface {
	Name() string
	Description() string
}
