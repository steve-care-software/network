package layers

import (
	"errors"

	"steve.care/network/domain/hash"
)

type engineBuilder struct {
	hashAdapter hash.Adapter
	execution   Execution
	resource    AssignableResource
}

func createEngineBuilder(
	hashAdapter hash.Adapter,
) EngineBuilder {
	out := engineBuilder{
		hashAdapter: hashAdapter,
		execution:   nil,
		resource:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *engineBuilder) Create() EngineBuilder {
	return createEngineBuilder(
		app.hashAdapter,
	)
}

// WithExecution adds an execution to the builder
func (app *engineBuilder) WithExecution(execution Execution) EngineBuilder {
	app.execution = execution
	return app
}

// WithResource adds a resource to the builder
func (app *engineBuilder) WithResource(resource AssignableResource) EngineBuilder {
	app.resource = resource
	return app
}

// Now builds a new Engine instance
func (app *engineBuilder) Now() (Engine, error) {
	data := [][]byte{}
	if app.execution != nil {
		data = append(data, []byte("execution"))
		data = append(data, app.execution.Hash().Bytes())
	}

	if app.resource != nil {
		data = append(data, []byte("resource"))
		data = append(data, app.resource.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Engine is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.execution != nil {
		return createEngineWithExecution(*pHash, app.execution), nil
	}

	return createEngineWithResource(*pHash, app.resource), nil
}
