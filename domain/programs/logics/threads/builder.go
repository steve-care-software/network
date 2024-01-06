package threads

import (
	"errors"

	"steve.care/network/domain/hash"
)

type builder struct {
	hashAapter hash.Adapter
	list       []Thread
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
func (app *builder) WithList(list []Thread) Builder {
	app.list = list
	return app
}

// Now builds a new Threads instance
func (app *builder) Now() (Threads, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Thread in order to build an Executiosn instance")
	}

	data := [][]byte{}
	for _, oneThread := range app.list {
		data = append(data, oneThread.Hash().Bytes())
	}

	pHash, err := app.hashAapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createThreads(*pHash, app.list), nil
}
