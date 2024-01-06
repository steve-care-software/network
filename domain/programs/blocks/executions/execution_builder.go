package executions

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/executions/actions"
)

type executionBuilder struct {
	hashAdapter hash.Adapter
	actions     actions.Actions
	receipt     hash.Hash
}

func createExecutionBuilder(
	hashAdapter hash.Adapter,
) ExecutionBuilder {
	out := executionBuilder{
		hashAdapter: hashAdapter,
		actions:     nil,
		receipt:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder(
		app.hashAdapter,
	)
}

// WithActions add actions to the builder
func (app *executionBuilder) WithActions(actions actions.Actions) ExecutionBuilder {
	app.actions = actions
	return app
}

// WithReceipt add receipt to the builder
func (app *executionBuilder) WithReceipt(receipt hash.Hash) ExecutionBuilder {
	app.receipt = receipt
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.actions == nil {
		return nil, errors.New("the actions is mandatory in order to build an Execution instance")
	}

	data := [][]byte{
		app.actions.Hash().Bytes(),
	}

	if app.receipt != nil {
		data = append(data, app.receipt.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.receipt != nil {
		return createExecutionWithReceipt(*pHash, app.actions, app.receipt), nil
	}

	return createExecution(*pHash, app.actions), nil
}
