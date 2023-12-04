package conditions

import "errors"

type pointerBuilder struct {
	entity string
	field  string
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		entity: "",
		field:  "",
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder()
}

// WithEntity adds a entity to the builder
func (app *pointerBuilder) WithEntity(entity string) PointerBuilder {
	app.entity = entity
	return app
}

// WithField adds a field to the builder
func (app *pointerBuilder) WithField(field string) PointerBuilder {
	app.field = field
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.entity == "" {
		return nil, errors.New("the entity is mandatory in order to build a Pointer instance")
	}

	if app.field == "" {
		return nil, errors.New("the field is mandatory in order to build a Pointer instance")
	}

	return createPointer(app.entity, app.field), nil
}
