package dashboards

import (
	"errors"

	"steve.care/network/domain/hash"
)

type builder struct {
	hashAapter hash.Adapter
	list       []Dashboard
}

func createBuilder(
	hashAapter hash.Adapter,
) Builder {
	out := builder{
		hashAapter: hashAapter,
		list:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAapter,
	)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Dashboard) Builder {
	app.list = list
	return app
}

// Now builds a new Dashboards instance
func (app *builder) Now() (Dashboards, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Dashboard in order to build an Executiosn instance")
	}

	data := [][]byte{}
	for _, oneDashboard := range app.list {
		data = append(data, oneDashboard.Hash().Bytes())
	}

	pHash, err := app.hashAapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createDashboards(*pHash, app.list), nil
}
