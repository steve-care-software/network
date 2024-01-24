package fields

import (
	"errors"
)

type fieldBuilder struct {
	name     string
	pKind    *uint8
	canBeNil bool
}

func createFieldBuilder() FieldBuilder {
	out := fieldBuilder{
		name:     "",
		pKind:    nil,
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

// WithKind adds a kind to the builder
func (app *fieldBuilder) WithKind(kind uint8) FieldBuilder {
	app.pKind = &kind
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

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Field instance")
	}

	return createField(app.name, *app.pKind, app.canBeNil), nil
}
