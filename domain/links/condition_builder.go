package links

import (
	"errors"

	"steve.care/network/domain/hash"
)

type conditionBuilder struct {
	hashAdapter hash.Adapter
	resource    ConditionResource
	operator    Operator
	next        ConditionValue
}

func createConditionBuilder(
	hashAdapter hash.Adapter,
) ConditionBuilder {
	out := conditionBuilder{
		hashAdapter: hashAdapter,
		resource:    nil,
		operator:    nil,
		next:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *conditionBuilder) Create() ConditionBuilder {
	return createConditionBuilder(
		app.hashAdapter,
	)
}

// WithResource adds a resource builder
func (app *conditionBuilder) WithResource(resource ConditionResource) ConditionBuilder {
	app.resource = resource
	return app
}

// WithOperator adds an operator builder
func (app *conditionBuilder) WithOperator(operator Operator) ConditionBuilder {
	app.operator = operator
	return app
}

// WithNext adds a next value to the builder builder
func (app *conditionBuilder) WithNext(next ConditionValue) ConditionBuilder {
	app.next = next
	return app
}

// Now builds a new Condition instance
func (app *conditionBuilder) Now() (Condition, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build an Condition instance")
	}

	if app.operator == nil {
		return nil, errors.New("the operator is mandatory in order to build an Condition instance")
	}

	if app.next == nil {
		return nil, errors.New("the next value is mandatory in order to build an Condition instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.resource.Hash().Bytes(),
		app.operator.Hash().Bytes(),
		app.next.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createCondition(*pHash, app.resource, app.operator, app.next), nil
}
