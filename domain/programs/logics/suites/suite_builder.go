package suites

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
	"steve.care/network/domain/programs/logics/suites/expectations"
)

type suiteBuilder struct {
	hashAdapter hash.Adapter
	origin      links.Origin
	input       []byte
	expectation expectations.Expectation
}

func createSuiteBuilder(
	hashAdapter hash.Adapter,
) SuiteBuilder {
	out := suiteBuilder{
		hashAdapter: hashAdapter,
		origin:      nil,
		input:       nil,
		expectation: nil,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder(
		app.hashAdapter,
	)
}

// WithOrigin adds an origin to the builder
func (app *suiteBuilder) WithOrigin(origin links.Origin) SuiteBuilder {
	app.origin = origin
	return app
}

// WithInput adds an input to the builder
func (app *suiteBuilder) WithInput(input []byte) SuiteBuilder {
	app.input = input
	return app
}

// WithExpectation adds an expectation to the builder
func (app *suiteBuilder) WithExpectation(expectation expectations.Expectation) SuiteBuilder {
	app.expectation = expectation
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.origin == nil {
		return nil, errors.New("the origin is mnadatory in order to build a Suite instance")
	}

	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.input == nil {
		return nil, errors.New("the input is mnadatory in order to build a Suite instance")
	}

	if app.expectation == nil {
		return nil, errors.New("the expectation is mnadatory in order to build a Suite instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.origin.Hash().Bytes(),
		app.input,
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
