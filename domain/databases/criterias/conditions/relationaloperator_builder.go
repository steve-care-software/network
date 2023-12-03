package conditions

import "errors"

type relationalOperatorBuilder struct {
	isAnd bool
	isOr  bool
}

func createRelationalOperatorBuilder() RelationalOperatorBuilder {
	out := relationalOperatorBuilder{
		isAnd: true,
		isOr:  true,
	}

	return &out
}

// Create initializes the builder
func (app *relationalOperatorBuilder) Create() RelationalOperatorBuilder {
	return createRelationalOperatorBuilder()
}

// IsAnd adds an and to the builder
func (app *relationalOperatorBuilder) IsAnd() RelationalOperatorBuilder {
	app.isAnd = true
	return app
}

// IsOr adds an or to the builder
func (app *relationalOperatorBuilder) IsOr() RelationalOperatorBuilder {
	app.isOr = true
	return app
}

// Now builds a new RelationalOperator instance
func (app *relationalOperatorBuilder) Now() (RelationalOperator, error) {
	if app.isAnd {
		return createRelationalOperatorWithAnd(), nil
	}

	if app.isOr {
		return createRelationalOperatorWithOr(), nil
	}

	return nil, errors.New("the RelationalOperator is invalid")
}
