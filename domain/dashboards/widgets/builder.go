package widgets

import (
	"errors"

	"steve.care/network/domain/hash"
)

type builder struct {
	hashAapter hash.Adapter
	list       []Widget
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
func (app *builder) WithList(list []Widget) Builder {
	app.list = list
	return app
}

// Now builds a new Widgets instance
func (app *builder) Now() (Widgets, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Widget in order to build an Executiosn instance")
	}

	data := [][]byte{}
	for _, oneWidget := range app.list {
		data = append(data, oneWidget.Hash().Bytes())
	}

	pHash, err := app.hashAapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createWidgets(*pHash, app.list), nil
}
