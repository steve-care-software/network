package layers

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/links"
)

type linkInstructionBuilder struct {
	hashAdapter hash.Adapter
	save        links.Link
	delete      hash.Hash
}

func createLinkInstructionBuilder(
	hashAdapter hash.Adapter,
) LinkInstructionBuilder {
	out := linkInstructionBuilder{
		hashAdapter: hashAdapter,
		save:        nil,
		delete:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *linkInstructionBuilder) Create() LinkInstructionBuilder {
	return createLinkInstructionBuilder(
		app.hashAdapter,
	)
}

// WithSave adds a save to the builder
func (app *linkInstructionBuilder) WithSave(save links.Link) LinkInstructionBuilder {
	app.save = save
	return app
}

// WithDelete adds a delete to the builder
func (app *linkInstructionBuilder) WithDelete(delete hash.Hash) LinkInstructionBuilder {
	app.delete = delete
	return app
}

// Now builds a new LinkInstruction
func (app *linkInstructionBuilder) Now() (LinkInstruction, error) {
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
		return nil, errors.New("the LinkInstruction is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.save != nil {
		return createLinkInstructionWithSave(*pHash, app.save), nil
	}

	return createLinkInstructionWithDelete(*pHash, app.delete), nil
}
