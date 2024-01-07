package viewports

import (
	"errors"
	"strconv"

	"steve.care/network/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	pRow        *uint
	height      uint
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pRow:        nil,
		height:      0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithRow adds a row to the builder
func (app *builder) WithRow(row uint) Builder {
	app.pRow = &row
	return app
}

// WithHeight adds a height to the builder
func (app *builder) WithHeight(height uint) Builder {
	app.height = height
	return app
}

// Now builds a new Viewport instance
func (app *builder) Now() (Viewport, error) {
	if app.pRow == nil {
		return nil, errors.New("the row is mandatory in order to build a Viewport instance")
	}

	if app.height <= 0 {
		return nil, errors.New("the height must be greater than zero (0) in order to build a Viewport instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(*app.pRow))),
		[]byte(strconv.Itoa(int(app.height))),
	})

	if err != nil {
		return nil, err
	}

	return createViewport(*pHash, *app.pRow, app.height), nil
}
