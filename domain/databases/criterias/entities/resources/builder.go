package resources

import (
	"errors"

	"steve.care/network/domain/databases/criterias/conditions"
)

type builder struct {
	container string
	condition conditions.Condition
}

func createBuilder() Builder {
	out := builder{
		container: "",
		condition: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithContainer adds a container to the builder
func (app *builder) WithContainer(container string) Builder {
	app.container = container
	return app
}

// WithCondition adds a condition to the builder
func (app *builder) WithCondition(condition conditions.Condition) Builder {
	app.condition = condition
	return app
}

// Now builds a new Resource instance
func (app *builder) Now() (Resource, error) {
	if app.container == "" {
		return nil, errors.New("the container is mandatory in order to build a Resource instance")
	}

	if app.condition == nil {
		return nil, errors.New("the condition is mandatory in order to build a Resource instance")
	}

	return createResource(
		app.container,
		app.condition,
	), nil
}
