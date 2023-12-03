package conditions

type integerOperatorBuilder struct {
	isSmallerThan bool
	isBiggerThan  bool
	isEqual       bool
}

func createIntegerOperatorBuilder() IntegerOperatorBuilder {
	out := integerOperatorBuilder{
		isSmallerThan: false,
		isBiggerThan:  false,
		isEqual:       false,
	}

	return &out
}

// Create initializes the builder
func (app *integerOperatorBuilder) Create() IntegerOperatorBuilder {
	return createIntegerOperatorBuilder()
}

// IsSmallerThan adds a smaller than to the builder
func (app *integerOperatorBuilder) IsSmallerThan() IntegerOperatorBuilder {
	app.isSmallerThan = true
	return app
}

// IsBiggerThan adds a bigger than to the builder
func (app *integerOperatorBuilder) IsBiggerThan() IntegerOperatorBuilder {
	app.isBiggerThan = true
	return app
}

// IsEqual adds an equal to the builder
func (app *integerOperatorBuilder) IsEqual() IntegerOperatorBuilder {
	app.isEqual = true
	return app
}

// Now builds a new IntegerOperator instance
func (app *integerOperatorBuilder) Now() (IntegerOperator, error) {
	if app.isSmallerThan && app.isEqual {
		return createIntegerOperatorWithSmallerThanAndEqual(), nil
	}

	if app.isSmallerThan {
		return createIntegerOperatorWithSmallerThan(), nil
	}

	if app.isBiggerThan && app.isEqual {
		return createIntegerOperatorWithBiggerThanAndEqual(), nil
	}

	if app.isBiggerThan {
		return createIntegerOperatorWithBiggerThan(), nil
	}

	if app.isEqual {
		return createIntegerOperatorWithEqual(), nil
	}

	return createIntegerOperator(), nil
}
