package suites

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
	"steve.care/network/domain/programs/logics/suites/expectations"
)

type builder struct {
	hashAdapter hash.Adapter
	origin      links.Origin
	input       layers.Layer
	expectation expectations.Expectation
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		origin:      nil,
		input:       nil,
		expectation: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithOrigin adds an origin to the builder
func (app *builder) WithOrigin(origin links.Origin) Builder {
	app.origin = origin
	return app
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input layers.Layer) Builder {
	app.input = input
	return app
}

// WithExpectation adds an expectation to the builder
func (app *builder) WithExpectation(expectation expectations.Expectation) Builder {
	app.expectation = expectation
	return app
}

// Now builds a new Suite instance
func (app *builder) Now() (Suite, error) {
	if app.origin == nil {
		return nil, errors.New("the origin is mnadatory in order to build a Suite instance")
	}

	if app.input == nil {
		return nil, errors.New("the input is mnadatory in order to build a Suite instance")
	}

	if app.expectation == nil {
		return nil, errors.New("the expectation is mnadatory in order to build a Suite instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.origin.Hash().Bytes(),
		app.input.Hash().Bytes(),
		app.expectation.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createSuite(
		*pHash,
		app.origin,
		app.input,
		app.expectation,
	), nil

}
