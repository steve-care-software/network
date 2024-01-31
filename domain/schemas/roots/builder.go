package roots

import (
	"errors"

	"steve.care/network/domain/schemas/roots/groups"
	"steve.care/network/domain/schemas/roots/methods"
)

type builder struct {
	name    string
	chains  groups.MethodChains
	methods methods.Methods
}

func createBuilder() Builder {
	out := builder{
		name:    "",
		chains:  nil,
		methods: nil,
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
func (app *builder) WithChains(chains groups.MethodChains) Builder {
	app.chains = chains
	return app
}

// WithMethods add methods to the builder
func (app *builder) WithMethods(methods methods.Methods) Builder {
	app.methods = methods
	return app
}

// Now builds a new Root instance
func (app *builder) Now() (Root, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatoryin order to build a Root instance")
	}

	if app.chains == nil {
		return nil, errors.New("the chains is mandatoryin order to build a Root instance")
	}

	if app.methods == nil {
		return nil, errors.New("the methods is mandatoryin order to build a Root instance")
	}

	return createRoot(app.name, app.chains, app.methods), nil
}
