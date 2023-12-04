package pointers

import "errors"

type builder struct {
	entity string
	field  string
}

func createBuilder() Builder {
	out := builder{
		entity: "",
		field:  "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithEntity adds an entity to the builder
func (app *builder) WithEntity(entity string) Builder {
	app.entity = entity
	return app
}

// WithField adds a field to the builder
func (app *builder) WithField(field string) Builder {
	app.field = field
	return app
}

// Now builds a new Pointer instance
func (app *builder) Now() (Pointer, error) {
	if app.entity == "" {
		return nil, errors.New("the entity is mandatory in order to build a Pointer instance")
	}

	if app.field == "" {
		return nil, errors.New("the field is mandatory in order to build a Pointer instance")
	}

	return createPointer(
		app.entity,
		app.field,
	), nil
}
