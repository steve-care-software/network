package schemas

import (
	"errors"

	"steve.care/network/domain/orms/schemas/roots"
)

type builder struct {
	roots    roots.Roots
	previous Schema
}

func createBuilder() Builder {
	out := builder{
		roots:    nil,
		previous: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRoots adds a roots to the builder
func (app *builder) WithRoots(roots roots.Roots) Builder {
	app.roots = roots
	return app
}

// WithPrevious add previous to the builder
func (app *builder) WithPrevious(previous Schema) Builder {
	app.previous = previous
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.roots == nil {
		return nil, errors.New("the roots is mandatory in order to build a Schema instance")
	}

	version := uint(0)
	if app.previous != nil {
		previousVersion := app.previous.Version()
		version = previousVersion + 1
	}

	return createSchema(version, app.roots), nil
}
