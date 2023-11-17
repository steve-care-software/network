package results

// Builder represents the result builder
type Builder interface {
	Create() Builder
	WithBytes(bytes []byte) Builder
	WithAction(action Action) Builder
	Now() (Result, error)
}

// Result represents result
type Result interface {
	Bytes() []byte
	Action() Action
}

// ActionBuilder represents the action builder
type ActionBuilder interface {
	Create() ActionBuilder
	IsPrompt() ActionBuilder
	IsContinue() ActionBuilder
	IsExecute() bool
	Execute() []string
}

// Action represents the action
type Action interface {
	IsPrompt() bool
	ISContinue() bool
	IsExecute() bool
	Execute() string
}
