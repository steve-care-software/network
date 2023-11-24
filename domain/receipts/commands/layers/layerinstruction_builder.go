package layers

import (
	"errors"

	"steve.care/network/domain/hash"
)

type layerInstructionBuilder struct {
	hashAdapter hash.Adapter
	save        Layer
	delete      hash.Hash
}

func createLayerInstructionBuilder(
	hashAdapter hash.Adapter,
) LayerInstructionBuilder {
	out := layerInstructionBuilder{
		hashAdapter: hashAdapter,
		save:        nil,
		delete:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *layerInstructionBuilder) Create() LayerInstructionBuilder {
	return createLayerInstructionBuilder(
		app.hashAdapter,
	)
}

// WithSave adds a save to the builder
func (app *layerInstructionBuilder) WithSave(save Layer) LayerInstructionBuilder {
	app.save = save
	return app
}

// WithDelete adds a delete to the builder
func (app *layerInstructionBuilder) WithDelete(delete hash.Hash) LayerInstructionBuilder {
	app.delete = delete
	return app
}

// Now builds a new LayerInstruction
func (app *layerInstructionBuilder) Now() (LayerInstruction, error) {
	data := [][]byte{}
	if app.save != nil {
		data = append(data, []byte("save"))
		data = append(data, app.save.Hash().Bytes())
	}

	if app.delete != nil {
		data = append(data, []byte("delete"))
		data = append(data, app.delete.Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the LayerInstruction is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.save != nil {
		return createLayerInstructionWithSave(*pHash, app.save), nil
	}

	return createLayerInstructionWithDelete(*pHash, app.delete), nil
}
