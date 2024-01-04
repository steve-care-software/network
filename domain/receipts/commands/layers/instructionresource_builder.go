package layers

import (
	"errors"

	"steve.care/network/domain/hash"
)

type instructionResourceBuilder struct {
	hashAdapter hash.Adapter
	save        string
	del         string
}

func createInstructionResourceBuilder(
	hashAdapter hash.Adapter,
) InstructionResourceBuilder {
	out := instructionResourceBuilder{
		hashAdapter: hashAdapter,
		save:        "",
		del:         "",
	}

	return &out
}

// Create initializes the builder
func (app *instructionResourceBuilder) Create() InstructionResourceBuilder {
	return createInstructionResourceBuilder(
		app.hashAdapter,
	)
}

// WithSave adds a save to the builder
func (app *instructionResourceBuilder) WithSave(save string) InstructionResourceBuilder {
	app.save = save
	return app
}

// WithDelete adds a delete to the builder
func (app *instructionResourceBuilder) WithDelete(del string) InstructionResourceBuilder {
	app.del = del
	return app
}

// Now builds a new InstructionResource instance
func (app *instructionResourceBuilder) Now() (InstructionResource, error) {
	data := [][]byte{}
	if app.save != "" {
		data = append(data, []byte("save"))
		data = append(data, []byte(app.save))
	}

	if app.del != "" {
		data = append(data, []byte("delete"))
		data = append(data, []byte(app.del))
	}

	if len(data) <= 0 {
		return nil, errors.New("the InstructionResource is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.save != "" {
		return createInstructionResourceWithSave(*pHash, app.save), nil
	}

	return createInstructionResourceWithDelete(*pHash, app.del), nil
}
