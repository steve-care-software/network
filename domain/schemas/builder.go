package schemas

import (
	"errors"

	"steve.care/network/domain/schemas/groups"
)

type builder struct {
	groups   groups.Groups
	previous Schema
}

func createBuilder() Builder {
	out := builder{
		groups:   nil,
		previous: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithGroups add groups to the builder
func (app *builder) WithGroups(groups groups.Groups) Builder {
	app.groups = groups
	return app
}

// WithPrevious add previous to the builder
func (app *builder) WithPrevious(previous Schema) Builder {
	app.previous = previous
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.groups == nil {
		return nil, errors.New("the groups is mandatory in order to build a Schema instance")
	}

	version := uint(0)
	if app.previous != nil {
		previousVersion := app.previous.Version()
		version = previousVersion + 1
	}

	return createSchema(version, app.groups), nil
}
