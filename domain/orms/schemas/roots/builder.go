package roots

import "errors"

type builder struct {
	list []Root
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
func (app *builder) WithList(list []Root) Builder {
	app.list = list
	return app
}

// Now builds a new Root instance
func (app *builder) Now() (Roots, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Root instance in order to build a Roots instance")
	}

	return createRoots(app.list), nil
}
