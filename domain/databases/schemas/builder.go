package schemas

import (
	"errors"

	"steve.care/network/domain/databases/schemas/connections"
	"steve.care/network/domain/databases/schemas/entities"
)

type builder struct {
	entities    entities.Entities
	connections connections.Connections
}

func createBuilder() Builder {
	out := builder{
		entities:    nil,
		connections: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithEntities add entities to the builder
func (app *builder) WithEntities(entities entities.Entities) Builder {
	app.entities = entities
	return app
}

// WithConnections add connections to the builder
func (app *builder) WithConnections(connections connections.Connections) Builder {
	app.connections = connections
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.entities == nil {
		return nil, errors.New("the entities is mandatory in order to build a Schema instance")
	}

	if app.connections != nil {
		return createSchemaWithConnections(
			app.entities,
			app.connections,
		), nil
	}

	return createSchema(
		app.entities,
	), nil
}
