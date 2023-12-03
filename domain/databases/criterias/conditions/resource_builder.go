package conditions

import (
	"errors"

	"steve.care/network/domain/databases/criterias/values"
)

type resourceBuilder struct {
	field Pointer
	value values.Value
}

func createResourceBuilder() ResourceBuilder {
	out := resourceBuilder{
		field: nil,
		value: nil,
	}

	return &out
}

// Create initializes the builder
func (app *resourceBuilder) Create() ResourceBuilder {
	return createResourceBuilder()
}

// WithField adds a field to the builder
func (app *resourceBuilder) WithField(field Pointer) ResourceBuilder {
	app.field = field
	return app
}

// WithValue adds a value to the builder
func (app *resourceBuilder) WithValue(value values.Value) ResourceBuilder {
	app.value = value
	return app
}

// Now builds a new Resource instance
func (app *resourceBuilder) Now() (Resource, error) {
	if app.field != nil {
		return createResourceWithField(app.field), nil
	}

	if app.value != nil {
		return createResourceWithValue(app.value), nil
	}

	return nil, errors.New("the Resource is invalid")
}
