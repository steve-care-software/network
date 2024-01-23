package resources

import (
	"errors"

	"steve.care/network/domain/schemas/resources/fields"
)

type resourceBuilder struct {
	name        string
	key         fields.Field
	fields      fields.Fields
	connections Connections
}

func createResourceBuilder() ResourceBuilder {
	out := resourceBuilder{
		name:        "",
		key:         nil,
		fields:      nil,
		connections: nil,
	}

	return &out
}

// Create initializes the builder
func (app *resourceBuilder) Create() ResourceBuilder {
	return createResourceBuilder()
}

// WithName adds a name to the builder
func (app *resourceBuilder) WithName(name string) ResourceBuilder {
	app.name = name
	return app
}

// WithKey adds a key to the builder
func (app *resourceBuilder) WithKey(key fields.Field) ResourceBuilder {
	app.key = key
	return app
}

// WithFields adds a fields to the builder
func (app *resourceBuilder) WithFields(fields fields.Fields) ResourceBuilder {
	app.fields = fields
	return app
}

// WithConnections adds a connections to the builder
func (app *resourceBuilder) WithConnections(connections Connections) ResourceBuilder {
	app.connections = connections
	return app
}

// Now builds a new Resource instance
func (app *resourceBuilder) Now() (Resource, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Resource instance")
	}

	if app.key == nil {
		return nil, errors.New("the key is mandatory in order to build a Resource instance")
	}

	if app.fields == nil {
		return nil, errors.New("the fields is mandatory in order to build a Resource instance")
	}

	if app.connections != nil {
		return createResourceWithConnections(app.name, app.key, app.fields, app.connections), nil
	}

	return createResource(app.name, app.key, app.fields), nil
}
