package expectations

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers"
	"steve.care/network/domain/receipts/commands/links"
)

type builder struct {
	hashAdapter hash.Adapter
	output      layers.Layer
	condition   links.Condition
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		output:      nil,
		condition:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithOutput adds an output to the builder
func (app *builder) WithOutput(output layers.Layer) Builder {
	app.output = output
	return app
}

// WithCondition adds a condition to the builder
func (app *builder) WithCondition(condition links.Condition) Builder {
	app.condition = condition
	return app
}

// Now builds a new Expectation instance
func (app *builder) Now() (Expectation, error) {
	data := [][]byte{}
	if app.output != nil {
		data = append(data, app.output.Hash().Bytes())
	}

	if app.condition != nil {
		data = append(data, app.output.Output().Hash().Bytes())
	}

	if len(data) != 1 {
		return nil, errors.New("the Expectation is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.output != nil {
		return createExpectationWithOutput(*pHash, app.output), nil
	}

	return createExpectationWithCondition(*pHash, app.condition), nil
}
