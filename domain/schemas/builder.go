package schemas

import (
	"errors"

	"steve.care/network/domain/schemas/resources"
)

type builder struct {
	resources resources.Resources
	previous  Schema
}

func createBuilder() Builder {
	out := builder{
		resources: nil,
		previous:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithResources add resources to the builder
func (app *builder) WithResources(resources resources.Resources) Builder {
	app.resources = resources
	return app
}

// WithPrevious add previous to the builder
func (app *builder) WithPrevious(previous Schema) Builder {
	app.previous = previous
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.resources == nil {
		return nil, errors.New("the resources is mandatory in order to build a Schema instance")
	}

	version := uint(0)
	if app.previous != nil {
		previousVersion := app.previous.Version()
		version = previousVersion + 1
	}

	return createSchema(version, app.resources), nil
}
