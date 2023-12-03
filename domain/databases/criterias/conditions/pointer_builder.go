package conditions

import "errors"

type pointerBuilder struct {
	container string
	field     string
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		container: "",
		field:     "",
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder()
}

// WithContainer adds a container to the builder
func (app *pointerBuilder) WithContainer(container string) PointerBuilder {
	app.container = container
	return app
}

// WithField adds a field to the builder
func (app *pointerBuilder) WithField(field string) PointerBuilder {
	app.field = field
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.container == "" {
		return nil, errors.New("the container is mandatory in order to build a Pointer instance")
	}

	if app.field == "" {
		return nil, errors.New("the field is mandatory in order to build a Pointer instance")
	}

	return createPointer(app.container, app.field), nil
}
