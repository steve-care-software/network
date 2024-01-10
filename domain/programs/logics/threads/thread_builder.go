package threads

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

type threadBuilder struct {
	hashAdapter hash.Adapter
	input       []byte
	entry       layers.Layer
}

func createThreadBuilder(
	hashAdapter hash.Adapter,
) ThreadBuilder {
	out := threadBuilder{
		hashAdapter: hashAdapter,
		input:       nil,
		entry:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *threadBuilder) Create() ThreadBuilder {
	return createThreadBuilder(
		app.hashAdapter,
	)
}

// WithInput adds an input to the builder
func (app *threadBuilder) WithInput(input []byte) ThreadBuilder {
	app.input = input
	return app
}

// WithEntry adds an entry to the builder
func (app *threadBuilder) WithEntry(entry layers.Layer) ThreadBuilder {
	app.entry = entry
	return app
}

// Now builds a new Thread instance
func (app *threadBuilder) Now() (Thread, error) {
	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build a Thread instance")
	}

	if app.entry == nil {
		return nil, errors.New("the entry is mandatory in order to build a Thread instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.input,
		app.entry.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createThread(*pHash, app.input, app.entry), nil
}
