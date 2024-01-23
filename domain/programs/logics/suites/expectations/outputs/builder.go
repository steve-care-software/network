package outputs

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

type builder struct {
	hashAdapter hash.Adapter
	kind        layers.Kind
	value       []byte
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		kind:        nil,
		value:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithKind adds a kind to the builder
func (app *builder) WithKind(kind layers.Kind) Builder {
	app.kind = kind
	return app
}

// WithValue adds a value to the builder
func (app *builder) WithValue(value []byte) Builder {
	app.value = value
	return app
}

// Now builds a new Output instance
func (app *builder) Now() (Output, error) {
	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build an Output instance")
	}

	if app.value != nil && len(app.value) <= 0 {
		app.value = nil
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build an Output instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.kind.Hash().Bytes(),
		app.value,
	})

	if err != nil {
		return nil, err
	}

	return createOutput(*pHash, app.kind, app.value), nil
}
