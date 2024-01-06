package blocks

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/executions"
)

type builder struct {
	hashAdapter hash.Adapter
	executions  executions.Executions
	parent      hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		executions:  nil,
		parent:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithExecutions add executions to the builder
func (app *builder) WithExecutions(executions executions.Executions) Builder {
	app.executions = executions
	return app
}

// WithParent add parent to the builder
func (app *builder) WithParent(parent hash.Hash) Builder {
	app.parent = parent
	return app
}

// Now builds a new Block instance
func (app *builder) Now() (Block, error) {
	if app.executions != nil {
		return nil, errors.New("the executions is mandatory in order to build a Block instance")
	}

	data := [][]byte{
		app.executions.Hash().Bytes(),
	}

	if app.parent != nil {
		data = append(data, app.parent.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		return createBlockWithParent(*pHash, app.executions, app.parent), nil
	}

	return createBlock(*pHash, app.executions), nil
}
