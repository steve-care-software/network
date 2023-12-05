package suites

import (
	"errors"

	"steve.care/network/domain/suites"
	"steve.care/network/domain/suites/expectations"
)

type builder struct {
	suite       suites.Suite
	expectation expectations.Expectation
}

func createBuilder() Builder {
	out := builder{
		suite:       nil,
		expectation: nil,
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

// Now builds a new Suite instance
func (app *builder) Now() (Suite, error) {
	if app.suite != nil {
		return createSuiteWithSuite(app.suite), nil
	}

	if app.expectation != nil {
		return createSuiteWithExpectation(app.expectation), nil
	}

	return nil, errors.New("the Suite is invalid")
}
