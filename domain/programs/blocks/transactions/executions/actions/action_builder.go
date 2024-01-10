package actions

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
)

type actionBuilder struct {
	hashAdapter hash.Adapter
	create      resources.Resource
	del         hash.Hash
}

func createActionBuilder(
	hashAdapter hash.Adapter,
) ActionBuilder {
	out := actionBuilder{
		hashAdapter: hashAdapter,
		create:      nil,
		del:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *actionBuilder) Create() ActionBuilder {
	return createActionBuilder(
		app.hashAdapter,
	)
}

// WithCreate adds a create to the builder
func (app *actionBuilder) WithCreate(create resources.Resource) ActionBuilder {
	app.create = create
	return app
}

// WithDelete adds a delete to the builder
func (app *actionBuilder) WithDelete(del hash.Hash) ActionBuilder {
	app.del = del
	return app
}

// Now builds a new Action instance
func (app *actionBuilder) Now() (Action, error) {
	data := [][]byte{}
	if app.create != nil {
		data = append(data, []byte("create"))
		data = append(data, app.create.Hash().Bytes())
	}

	if app.del != nil {
		data = append(data, []byte("delete"))
		data = append(data, app.del.Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Action is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.create != nil {
		return createActionWithCreate(*pHash, app.create), nil
	}

	return createActionWithDelete(*pHash, app.del), nil
}
