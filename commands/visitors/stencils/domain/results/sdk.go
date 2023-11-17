package results

// Result represents result
type Result interface {
	Bytes() []byte
	Action() Action
}

// Action represents the action
type Action interface {
	IsPrompt() bool
	ISContinue() bool
	IsExecute() bool
	Execute() string
}
