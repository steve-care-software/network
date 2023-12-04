package resources

import (
	"errors"

	"steve.care/network/domain/databases/criterias/conditions"
)

type builder struct {
	entity    string
	condition conditions.Condition
}

func createBuilder() Builder {
	out := builder{
		entity:    "",
		condition: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithEntity adds a entity to the builder
func (app *builder) WithEntity(entity string) Builder {
	app.entity = entity
	return app
}

// WithCondition adds a condition to the builder
func (app *builder) WithCondition(condition conditions.Condition) Builder {
	app.condition = condition
	return app
}

// Now builds a new Resource instance
func (app *builder) Now() (Resource, error) {
	if app.entity == "" {
		return nil, errors.New("the entity is mandatory in order to build a Resource instance")
	}

	if app.condition == nil {
		return nil, errors.New("the condition is mandatory in order to build a Resource instance")
	}

	return createResource(
		app.entity,
		app.condition,
	), nil
}
