package fields

import "errors"

type builder struct {
	list []Field
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

// WithList adds a list to the builder
func (app *builder) WithList(list []Field) Builder {
	app.list = list
	return app
}

// Now builds a new Fields instance
func (app *builder) Now() (Fields, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Field in order to build a Fields instance")
	}

	return createFields(app.list), nil
}
