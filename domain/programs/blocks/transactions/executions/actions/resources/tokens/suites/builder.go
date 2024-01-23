package suites

import (
	"errors"

	"steve.care/network/domain/programs/logics/suites"
	"steve.care/network/domain/programs/logics/suites/expectations"
	"steve.care/network/domain/programs/logics/suites/expectations/outputs"
)

type builder struct {
	suite       suites.Suite
	expectation expectations.Expectation
	output      outputs.Output
}

func createBuilder() Builder {
	out := builder{
		suite:       nil,
		expectation: nil,
		output:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithSuite adds a suite to the builder
func (app *builder) WithSuite(suite suites.Suite) Builder {
	app.suite = suite
	return app
}

// WithExpectation adds an expectation to the builder
func (app *builder) WithExpectation(expectation expectations.Expectation) Builder {
	app.expectation = expectation
	return app
}

// WithOutput adds an output to the builder
func (app *builder) WithOutput(output outputs.Output) Builder {
	app.output = output
	return app
}

// Now builds a new Suite instance
func (app *builder) Now() (Suite, error) {
	if app.suite != nil {
		return createSuiteWithSuite(app.suite), nil
	}

	if app.expectation != nil {
		return createSuiteWithExpectation(app.expectation), nil
	}

	if app.output != nil {
		return createSuiteWithOutput(app.output), nil
	}

	return nil, errors.New("the Suite is invalid")
}
