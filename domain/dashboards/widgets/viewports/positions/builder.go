package positions

import (
	"errors"
	"fmt"

	"steve.care/network/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	horizontal  float32
	vertical    float32
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		horizontal:  1,
		vertical:    1,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithHorizontal adds a horizontal to the builder
func (app *builder) WithHorizontal(horizontal float32) Builder {
	app.horizontal = horizontal
	return app
}

// WithVertical adds an vertical to the builder
func (app *builder) WithVertical(vertical float32) Builder {
	app.vertical = vertical
	return app
}

// Now builds a new Position instance
func (app *builder) Now() (Position, error) {
	if app.horizontal >= 1 {
		return nil, errors.New("the horizontal must be smaller than one (1) in order to byuild a Position instance")
	}

	if app.vertical >= 1 {
		return nil, errors.New("the vertical must be smaller than one (1) in order to byuild a Position instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(fmt.Sprintf("%f", app.horizontal)),
		[]byte(fmt.Sprintf("%f", app.vertical)),
	})

	if err != nil {
		return nil, err
	}

	return createPosition(*pHash, app.horizontal, app.vertical), nil
}
