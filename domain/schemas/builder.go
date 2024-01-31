package schemas

import (
	"errors"

	"steve.care/network/domain/schemas/roots"
)

type builder struct {
	root     roots.Root
	previous Schema
}

func createBuilder() Builder {
	out := builder{
		root:     nil,
		previous: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root roots.Root) Builder {
	app.root = root
	return app
}

// WithPrevious add previous to the builder
func (app *builder) WithPrevious(previous Schema) Builder {
	app.previous = previous
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.root == nil {
		return nil, errors.New("the root is mandatory in order to build a Schema instance")
	}

	version := uint(0)
	if app.previous != nil {
		previousVersion := app.previous.Version()
		version = previousVersion + 1
	}

	return createSchema(version, app.root), nil
}
