package stacks

type factory struct {
}

func createFactory() Factory {
	out := factory{}
	return &out
}

// Create creates a new empty stack
func (app *factory) Create() Stack {
	return createStack()
}
