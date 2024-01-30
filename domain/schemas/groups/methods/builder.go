package methods

import (
	"errors"

	resource_methods "steve.care/network/domain/schemas/groups/resources/methods"
)

type builder struct {
	builder resource_methods.Methods
}

func createBuilder() Builder {
	out := builder{
		builder: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithBuilder adds a builder to the builder
func (app *builder) WithBuilder(builder resource_methods.Methods) Builder {
	app.builder = builder
	return app
}

// WithBuilder adds a builder to the builder
func (app *builder) Now() (Methods, error) {
	if app.builder == nil {
		return nil, errors.New("the builder methods is mandatory in order to build a Methods instance")
	}

	return createMethods(app.builder), nil
}
