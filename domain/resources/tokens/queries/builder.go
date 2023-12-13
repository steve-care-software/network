package queries

import (
	"errors"

	"steve.care/network/domain/criterias"
	"steve.care/network/domain/criterias/conditions"
)

type builder struct {
	criteria    criterias.Criteria
	condition   conditions.Condition
	pointer     conditions.Pointer
	element     conditions.Element
	resource    conditions.Resource
	operator    conditions.Operator
	relOperator conditions.RelationalOperator
	intOperator conditions.IntegerOperator
}

func createBuilder() Builder {
	out := builder{
		criteria:    nil,
		condition:   nil,
		pointer:     nil,
		element:     nil,
		resource:    nil,
		operator:    nil,
		relOperator: nil,
		intOperator: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithCriteria adds a criteria to the builder
func (app *builder) WithCriteria(criteria criterias.Criteria) Builder {
	app.criteria = criteria
	return app
}

// WithCondition adds a condition to the builder
func (app *builder) WithCondition(condition conditions.Condition) Builder {
	app.condition = condition
	return app
}

// WithPointer adds a pointer to the builder
func (app *builder) WithPointer(pointer conditions.Pointer) Builder {
	app.pointer = pointer
	return app
}

// WithElement adds an element to the builder
func (app *builder) WithElement(element conditions.Element) Builder {
	app.element = element
	return app
}

// WithResource adds a resource to the builder
func (app *builder) WithResource(resource conditions.Resource) Builder {
	app.resource = resource
	return app
}

// WithOperator adds an operator to the builder
func (app *builder) WithOperator(operator conditions.Operator) Builder {
	app.operator = operator
	return app
}

// WithRelationalOperator adds a relational operator to the builder
func (app *builder) WithRelationalOperator(relOperator conditions.RelationalOperator) Builder {
	app.relOperator = relOperator
	return app
}

// WithIntegerOperator adds an integer operator to the builder
func (app *builder) WithIntegerOperator(intOperator conditions.IntegerOperator) Builder {
	app.intOperator = intOperator
	return app
}

// Now builds a new Query instance
func (app *builder) Now() (Query, error) {
	if app.criteria != nil {
		return createQueryWithCriteria(app.criteria), nil
	}

	if app.condition != nil {
		return createQueryWithCondition(app.condition), nil
	}

	if app.pointer != nil {
		return createQueryWithPointer(app.pointer), nil
	}

	if app.element != nil {
		return createQueryWithElement(app.element), nil
	}

	if app.resource != nil {
		return createQueryWithResource(app.resource), nil
	}

	if app.operator != nil {
		return createQueryWithOperator(app.operator), nil
	}

	if app.relOperator != nil {
		return createQueryWithRelationalOperator(app.relOperator), nil
	}

	if app.intOperator != nil {
		return createQueryWithIntegerOperator(app.intOperator), nil
	}

	return nil, errors.New("the Query is invalid")
}
