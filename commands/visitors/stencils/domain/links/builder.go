package links

import (
	"errors"

	"steve.care/network/libraries/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	origin      Origin
	elements    Elements
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		origin:      nil,
		elements:    nil,
	}

	return &out
}

// Builder initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithOrigin adds an origin to the builder
func (app *builder) WithOrigin(origin Origin) Builder {
	app.origin = origin
	return app
}

// WithElements add elements to the builder
func (app *builder) WithElements(elements Elements) Builder {
	app.elements = elements
	return app
}

// Now builds a new Link instance
func (app *builder) Now() (Link, error) {
	if app.origin == nil {
		return nil, errors.New("the origin is mandatory in order to build a Link instance")
	}

	if app.elements == nil {
		return nil, errors.New("the elements is mandatory in order to build a Link instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.origin.Hash().Bytes(),
		app.elements.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLink(*pHash, app.origin, app.elements), nil
}
