package methods

import (
	"errors"
)

type builder struct {
	initialize string
	trigger    string
	builder    string
}

func createBuilder() Builder {
	out := builder{
		initialize: "",
		trigger:    "",
		builder:    "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithInitialize adds an initialize to the builder
func (app *builder) WithInitialize(initialize string) Builder {
	app.initialize = initialize
	return app
}

// WithTrigger adds a trigger to the builder
func (app *builder) WithTrigger(trigger string) Builder {
	app.trigger = trigger
	return app
}

// WithBuilder adds a builder to the builder
func (app *builder) WithBuilder(builder string) Builder {
	app.builder = builder
	return app
}

// Now builds a new Methods instance
func (app *builder) Now() (Methods, error) {
	if app.initialize == "" {
		return nil, errors.New("the initialize is mandatory in order to build a Methods instance")
	}

	if app.trigger == "" {
		return nil, errors.New("the trigger is mandatory in order to build a Methods instance")
	}

	if app.builder == "" {
		return nil, errors.New("the builder is mandatory in order to build a Methods instance")
	}

	return createMethods(app.initialize, app.trigger, app.builder), nil
}
