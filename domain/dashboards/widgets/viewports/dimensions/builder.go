package dimensions

import (
	"errors"
	"fmt"

	"steve.care/network/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	width       float32
	height      float32
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		width:       0,
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

// WithWidth adds a width to the builder
func (app *builder) WithWidth(width float32) Builder {
	app.width = width
	return app
}

// WithHeight adds an height to the builder
func (app *builder) WithHeight(height float32) Builder {
	app.height = height
	return app
}

// Now builds a new Dimension instance
func (app *builder) Now() (Dimension, error) {
	if app.width <= 0 {
		return nil, errors.New("the width must be greater than zero (0) in order to byuild a Dimension instance")
	}

	if app.height <= 0 {
		return nil, errors.New("the height must be greater than zero (0) in order to byuild a Dimension instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(fmt.Sprintf("%f", app.width)),
		[]byte(fmt.Sprintf("%f", app.height)),
	})

	if err != nil {
		return nil, err
	}

	return createDimension(*pHash, app.width, app.height), nil
}
