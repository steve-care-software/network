package viewports

import (
	"errors"
	"strconv"

	"steve.care/network/domain/dashboards/widgets/viewports/dimensions"
	"steve.care/network/domain/dashboards/widgets/viewports/positions"
	"steve.care/network/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	pLayer      *uint
	position    positions.Position
	dimension   dimensions.Dimension
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pLayer:      nil,
		position:    nil,
		dimension:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithLayer adds a layer to the builder
func (app *builder) WithLayer(layer uint) Builder {
	app.pLayer = &layer
	return app
}

// WithPosition adds a position to the builder
func (app *builder) WithPosition(position positions.Position) Builder {
	app.position = position
	return app
}

// WithDimension adds a dimension to the builder
func (app *builder) WithDimension(dimension dimensions.Dimension) Builder {
	app.dimension = dimension
	return app
}

// Now builds a new Viewport instance
func (app *builder) Now() (Viewport, error) {
	if app.pLayer == nil {
		return nil, errors.New("the layer is mandatory in order to build a Viewport instance")
	}

	if app.position == nil {
		return nil, errors.New("the position is mandatory in order to build a Viewport instance")
	}

	if app.dimension == nil {
		return nil, errors.New("the dimension is mandatory in order to build a Viewport instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(*app.pLayer))),
		app.position.Hash().Bytes(),
		app.dimension.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createViewport(*pHash, *app.pLayer, app.position, app.dimension), nil
}
