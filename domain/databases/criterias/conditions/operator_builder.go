package conditions

import "errors"

type operatorBuilder struct {
	isEqual    bool
	relational RelationalOperator
	integer    IntegerOperator
}

func createOperatorBuilder() OperatorBuilder {
	out := operatorBuilder{
		isEqual:    false,
		relational: nil,
		integer:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *operatorBuilder) Create() OperatorBuilder {
	return createOperatorBuilder()
}

// WithRelational adds a relational to the builder
func (app *operatorBuilder) WithRelational(relational RelationalOperator) OperatorBuilder {
	app.relational = relational
	return app
}

// WithInteger adds an integer to the builder
func (app *operatorBuilder) WithInteger(integer IntegerOperator) OperatorBuilder {
	app.integer = integer
	return app
}

// IsEqual flags the builder as equal
func (app *operatorBuilder) IsEqual() OperatorBuilder {
	app.isEqual = true
	return app
}

// Now builds a new Operator instance
func (app *operatorBuilder) Now() (Operator, error) {
	if app.isEqual {
		return createOperatorWithEqual(), nil
	}

	if app.relational != nil {
		return createOperatorWithRelational(app.relational), nil
	}

	if app.integer != nil {
		return createOperatorWithInteger(app.integer), nil
	}

	return nil, errors.New("the Operator is invalid")
}
