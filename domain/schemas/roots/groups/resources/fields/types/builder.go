package types

import (
	"errors"

	"steve.care/network/domain/schemas/roots/groups/resources/fields/types/dependencies"
)

type builder struct {
	pKind      *uint8
	dependency dependencies.Dependency
}

func createBuilder() Builder {
	out := builder{
		pKind:      nil,
		dependency: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithKind adds a kind to the builder
func (app *builder) WithKind(kind uint8) Builder {
	app.pKind = &kind
	return app
}

// WithDependency adds a dependency to the builder
func (app *builder) WithDependency(dependency dependencies.Dependency) Builder {
	app.dependency = dependency
	return app
}

// Now builds a new Type instance
func (app *builder) Now() (Type, error) {
	if app.pKind != nil {
		return createTypeWithKind(app.pKind), nil
	}

	if app.dependency != nil {
		return createTypeWithDependency(app.dependency), nil
	}

	return nil, errors.New("the Type is invalid")
}
