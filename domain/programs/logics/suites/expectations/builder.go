package expectations

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
	"steve.care/network/domain/programs/logics/suites/expectations/outputs"
)

type builder struct {
	hashAdapter hash.Adapter
	success     outputs.Output
	mistake     links.Condition
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		success:     nil,
		mistake:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithSuccess adds a success to the builder
func (app *builder) WithSuccess(success outputs.Output) Builder {
	app.success = success
	return app
}

// WithMistake adds a mistake to the builder
func (app *builder) WithMistake(mistake links.Condition) Builder {
	app.mistake = mistake
	return app
}

// Now builds a new Expectation instance
func (app *builder) Now() (Expectation, error) {
	data := [][]byte{}
	if app.success != nil {
		data = append(data, app.success.Hash().Bytes())
	}

	if app.mistake != nil {
		data = append(data, app.mistake.Hash().Bytes())
	}

	if len(data) != 1 {
		return nil, errors.New("the Expectation is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.success != nil {
		return createExpectationWithSuccess(*pHash, app.success), nil
	}

	return createExpectationWithMistake(*pHash, app.mistake), nil
}
