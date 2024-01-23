package suites

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/suites"
	"steve.care/network/domain/programs/logics/suites/expectations"
	"steve.care/network/domain/programs/logics/suites/expectations/outputs"
)

type suite struct {
	suite       suites.Suite
	expectation expectations.Expectation
	output      outputs.Output
}

func createSuiteWithSuite(
	suiteIns suites.Suite,
) Suite {
	return createSuiteInternally(
		suiteIns,
		nil,
		nil,
	)
}

func createSuiteWithExpectation(
	expectation expectations.Expectation,
) Suite {
	return createSuiteInternally(
		nil,
		expectation,
		nil,
	)
}

func createSuiteWithOutput(
	output outputs.Output,
) Suite {
	return createSuiteInternally(
		nil,
		nil,
		output,
	)
}

func createSuiteInternally(
	suiteIns suites.Suite,
	expectation expectations.Expectation,
	output outputs.Output,
) Suite {
	out := suite{
		suite:       suiteIns,
		expectation: expectation,
		output:      output,
	}

	return &out
}

// Hash returns the hash
func (obj *suite) Hash() hash.Hash {
	if obj.IsSuite() {
		return obj.suite.Hash()
	}

	if obj.IsExpectation() {
		return obj.expectation.Hash()
	}

	return obj.output.Hash()
}

// IsSuite returns true if there is a suite, false otherwise
func (obj *suite) IsSuite() bool {
	return obj.suite != nil
}

// Suite returns the suite, if any
func (obj *suite) Suite() suites.Suite {
	return obj.suite
}

// IsExpectation returns true if there is an expectation, false otherwise
func (obj *suite) IsExpectation() bool {
	return obj.expectation != nil
}

// Expectation returns the expectation, if any
func (obj *suite) Expectation() expectations.Expectation {
	return obj.expectation
}

// IsOutput returns true if there is an output, false otherwise
func (obj *suite) IsOutput() bool {
	return obj.output != nil
}

// Output returns the output, if any
func (obj *suite) Output() outputs.Output {
	return obj.output
}
