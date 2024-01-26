package groups

import "errors"

type builder struct {
	name   string
	chains MethodChains
}

func createBuilder() Builder {
	out := builder{
		name:   "",
		chains: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithChains add chains to the builder
func (app *builder) WithChains(chains MethodChains) Builder {
	app.chains = chains
	return app
}

// Now builds a new Group instance
func (app *builder) Now() (Group, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Group instance")
	}

	if app.chains == nil {
		return nil, errors.New("the chains is mandatory in order to build a Group instance")
	}

	return createGroup(app.name, app.chains), nil
}
