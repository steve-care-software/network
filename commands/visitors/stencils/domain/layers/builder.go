package layers

import (
	"errors"
	"strings"

	"steve.care/network/libraries/hash"
)

type builder struct {
	hashAdapter  hash.Adapter
	path         []string
	input        string
	instructions Instructions
	output       Output
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		path:         nil,
		input:        "",
		instructions: nil,
		output:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path []string) Builder {
	app.path = path
	return app
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input string) Builder {
	app.input = input
	return app
}

// WithInstructions add instructions to the builder
func (app *builder) WithInstructions(instructions Instructions) Builder {
	app.instructions = instructions
	return app
}

// WithOutput add output to the builder
func (app *builder) WithOutput(output Output) Builder {
	app.output = output
	return app
}

// Now builds a new Layer instance
func (app *builder) Now() (Layer, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Layer instance")
	}

	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to build a Layer instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Layer instance")
	}

	if app.output == nil {
		return nil, errors.New("the output is mandatory in order to build a Layer instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strings.Join(app.path, "/")),
		[]byte(app.input),
		app.instructions.Hash().Bytes(),
		app.output.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLayer(*pHash, app.path, app.input, app.instructions, app.output), nil

}
