package entities

import (
	"errors"

	"steve.care/network/domain/databases/schemas/entities/fields"
)

type entityBuilder struct {
	name   string
	head   string
	fields fields.Fields
}

func createEntityBuilder() EntityBuilder {
	out := entityBuilder{
		name:   "",
		head:   "",
		fields: nil,
	}

	return &out
}

// Create initializes the builder
func (app *entityBuilder) Create() EntityBuilder {
	return createEntityBuilder()
}

// WithName adds a name to the builder
func (app *entityBuilder) WithName(name string) EntityBuilder {
	app.name = name
	return app
}

// WithHead adds an head to the builder
func (app *entityBuilder) WithHead(head string) EntityBuilder {
	app.head = head
	return app
}

// WithFields add fields to the builder
func (app *entityBuilder) WithFields(fields fields.Fields) EntityBuilder {
	app.fields = fields
	return app
}

// Now builds a new Entity instance
func (app *entityBuilder) Now() (Entity, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Entity instance")
	}

	if app.head == "" {
		return nil, errors.New("the head is mandatory in order to build an Entity instance")
	}

	if app.fields == nil {
		return nil, errors.New("the fields is mandatory in order to build an Entity instance")
	}

	return createEntity(
		app.name,
		app.head,
		app.fields,
	), nil
}
