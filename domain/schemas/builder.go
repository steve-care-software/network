package schemas

import (
	"errors"

	"steve.care/network/domain/schemas/groups"
)

type builder struct {
	group    groups.Group
	previous Schema
}

func createBuilder() Builder {
	out := builder{
		group:    nil,
		previous: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithGroup adds a group to the builder
func (app *builder) WithGroup(group groups.Group) Builder {
	app.group = group
	return app
}

// WithPrevious add previous to the builder
func (app *builder) WithPrevious(previous Schema) Builder {
	app.previous = previous
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.group == nil {
		return nil, errors.New("the group is mandatory in order to build a Schema instance")
	}

	version := uint(0)
	if app.previous != nil {
		previousVersion := app.previous.Version()
		version = previousVersion + 1
	}

	return createSchema(version, app.group), nil
}
