package fields

import (
	"errors"

	"steve.care/network/domain/orms/schemas/roots/groups/resources/fields/methods"
	"steve.care/network/domain/orms/schemas/roots/groups/resources/fields/types"
)

type fieldBuilder struct {
	name     string
	methods  methods.Methods
	typ      types.Type
	canBeNil bool
}

func createFieldBuilder() FieldBuilder {
	out := fieldBuilder{
		name:     "",
		methods:  nil,
		typ:      nil,
		canBeNil: false,
	}

	return &out
}

// Create initializes the builder
func (app *fieldBuilder) Create() FieldBuilder {
	return createFieldBuilder()
}

// WithName adds a name to the builder
func (app *fieldBuilder) WithName(name string) FieldBuilder {
	app.name = name
	return app
}

// WithMethods add methods to the builder
func (app *fieldBuilder) WithMethods(methods methods.Methods) FieldBuilder {
	app.methods = methods
	return app
}

// WithType adds a type to the builder
func (app *fieldBuilder) WithType(typ types.Type) FieldBuilder {
	app.typ = typ
	return app
}

// CanBeNil flags the builder as canBeNil
func (app *fieldBuilder) CanBeNil() FieldBuilder {
	app.canBeNil = true
	return app
}

// Now builds a new Field instance
func (app *fieldBuilder) Now() (Field, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Field instance")
	}

	if app.methods == nil {
		return nil, errors.New("the methods is mandatory in order to build a Field instance")
	}

	if app.typ == nil {
		return nil, errors.New("the type is mandatory in order to build a Field instance")
	}

	return createField(
		app.name,
		app.methods,
		app.typ,
		app.canBeNil,
	), nil
}
