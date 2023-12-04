package fields

import (
	"errors"

	"steve.care/network/domain/databases/schemas/entities/fields/kinds"
)

type fieldBuilder struct {
	name     string
	kind     kinds.Kind
	isUnique bool
}

func createFieldBuilder() FieldBuilder {
	out := fieldBuilder{
		name:     "",
		kind:     nil,
		isUnique: false,
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
func (app *fieldBuilder) WithKind(kind kinds.Kind) FieldBuilder {
	app.kind = kind
	return app
}

// IsUnique flags the builder as unique
func (app *fieldBuilder) IsUnique() FieldBuilder {
	app.isUnique = true
	return app
}

// Now builds a new Field instance
func (app *fieldBuilder) Now() (Field, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Field instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Field instance")
	}

	return createField(
		app.name,
		app.kind,
		app.isUnique,
	), nil
}
