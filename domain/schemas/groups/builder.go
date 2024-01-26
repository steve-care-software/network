package groups

import "errors"

type builder struct {
	list []Group
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
func (app *builder) WithList(list []Group) Builder {
	app.list = list
	return app
}

// Now builds a new Groups instance
func (app *builder) Now() (Groups, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Group in order to build a Groups instance")
	}

	mp := map[string]Group{}
	for _, oneGroup := range app.list {
		name := oneGroup.Name()
		mp[name] = oneGroup
	}

	return createGroups(mp, app.list), nil
}
