package groups

import "errors"

type methodChainsBuilder struct {
	list []MethodChain
}

func createMethodChainsBuilder() MethodChainsBuilder {
	out := methodChainsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *methodChainsBuilder) Create() MethodChainsBuilder {
	return createMethodChainsBuilder()
}

// WithList adds a list to the builder
func (app *methodChainsBuilder) WithList(list []MethodChain) MethodChainsBuilder {
	app.list = list
	return app
}

// Now builds a new MethodChains instance
func (app *methodChainsBuilder) Now() (MethodChains, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 MethodChain in order to build a MethodChains instance")
	}

	return createMethodChains(app.list), nil
}
