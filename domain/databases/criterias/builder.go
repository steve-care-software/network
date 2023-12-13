package criterias

import (
	"errors"

	"steve.care/network/domain/databases/criterias/conditions"
	"steve.care/network/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	entity      string
	condition   conditions.Condition
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		entity:      "",
		condition:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
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

// Now builds a new Criteria instance
func (app *builder) Now() (Criteria, error) {
	if app.entity == "" {
		return nil, errors.New("the entity is mandatory in order to build a Criteria instance")
	}

	if app.condition == nil {
		return nil, errors.New("the condition is mandatory in order to build a Criteria instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.entity),
		app.condition.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createCriteria(
		*pHash,
		app.entity,
		app.condition,
	), nil
}
