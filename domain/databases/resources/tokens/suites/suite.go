package suites

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/suites"
	"steve.care/network/domain/suites/expectations"
)

type suite struct {
	suite       suites.Suite
	expectation expectations.Expectation
}

func createSuiteWithSuite(
	suiteIns suites.Suite,
) Suite {
	return createSuiteInternally(
		suiteIns,
		nil,
	)
}

func createSuiteWithExpectation(
	expectation expectations.Expectation,
) Suite {
	return createSuiteInternally(
		nil,
		expectation,
	)
}

func createSuiteInternally(
	suiteIns suites.Suite,
	expectation expectations.Expectation,
) Suite {
	out := suite{
		suite:       suiteIns,
		expectation: expectation,
	}

	return &out
}

// Hash returns the hash
func (obj *suite) Hash() hash.Hash {
	if obj.IsSuite() {
		return obj.suite.Hash()
	}

	return obj.expectation.Hash()
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
