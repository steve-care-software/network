package roots

import (
	"errors"

	"steve.care/network/domain/orms/schemas/roots/groups"
	"steve.care/network/domain/orms/schemas/roots/methods"
)

type rootBuilder struct {
	name    string
	chains  groups.MethodChains
	methods methods.Methods
}

func createRootBuilder() RootBuilder {
	out := rootBuilder{
		name:    "",
		chains:  nil,
		methods: nil,
	}

	return &out
}

// Create initializes the builder
func (app *rootBuilder) Create() RootBuilder {
	return createRootBuilder()
}

// WithName adds a name to the builder
func (app *rootBuilder) WithName(name string) RootBuilder {
	app.name = name
	return app
}

// WithChains add chains to the builder
func (app *rootBuilder) WithChains(chains groups.MethodChains) RootBuilder {
	app.chains = chains
	return app
}

// WithMethods add methods to the builder
func (app *rootBuilder) WithMethods(methods methods.Methods) RootBuilder {
	app.methods = methods
	return app
}

// Now builds a new Root instance
func (app *rootBuilder) Now() (Root, error) {
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
