package resources

import "errors"

type pointerBuilder struct {
	resource Resource
	field    string
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		resource: nil,
		field:    "",
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder()
}

// WithResource adds a resource to the builder
func (app *pointerBuilder) WithResource(resource Resource) PointerBuilder {
	app.resource = resource
	return app
}

// WithField adds a field to the builder
func (app *pointerBuilder) WithField(field string) PointerBuilder {
	app.field = field
	return app
}

// Now builds a new Pointer
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build a Pointer instance")
	}

	if app.field == "" {
		return nil, errors.New("the field is mandatory in order to build a Pointer instance")
	}

	return createPointer(app.resource, app.field), nil
}
