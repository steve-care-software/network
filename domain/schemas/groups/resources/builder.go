package resources

import (
	"errors"

	"steve.care/network/domain/schemas/groups/resources/fields"
)

type builder struct {
	name        string
	key         fields.Field
	fields      fields.Fields
	connections Connections
}

func createBuilder() Builder {
	out := builder{
		name:        "",
		key:         nil,
		fields:      nil,
		connections: nil,
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

// WithKey adds a key to the builder
func (app *builder) WithKey(key fields.Field) Builder {
	app.key = key
	return app
}

// WithFields adds a fields to the builder
func (app *builder) WithFields(fields fields.Fields) Builder {
	app.fields = fields
	return app
}

// WithConnections adds a connections to the builder
func (app *builder) WithConnections(connections Connections) Builder {
	app.connections = connections
	return app
}

// Now builds a new Resource instance
func (app *builder) Now() (Resource, error) {
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
