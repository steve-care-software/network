package entities

import "errors"

type builder struct {
	list []Entity
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add list of entities to the builder
func (app *builder) WithList(list []Entity) Builder {
	app.list = list
	return app
}

// Now builds a new Entities instance
func (app *builder) Now() (Entities, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Entity in order to build a Entities instance")
	}

	return createEntities(
		app.list,
	), nil
}
