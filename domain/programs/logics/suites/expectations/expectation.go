package expectations

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
	"steve.care/network/domain/programs/logics/suites/expectations/outputs"
)

type expectation struct {
	hash    hash.Hash
	success outputs.Output
	mistake links.Condition
}

func createExpectationWithSuccess(
	hash hash.Hash,
	success outputs.Output,
) Expectation {
	return createExpectationInternally(
		hash,
		success,
		nil,
	)
}

func createExpectationWithMistake(
	hash hash.Hash,
	mistake links.Condition,
) Expectation {
	return createExpectationInternally(
		hash,
		nil,
		mistake,
	)
}

func createExpectationInternally(
	hash hash.Hash,
	success outputs.Output,
	mistake links.Condition,
) Expectation {
	out := expectation{
		hash:    hash,
		success: success,
		mistake: mistake,
	}

	return &out
}

// Hash returns the hash
func (obj *expectation) Hash() hash.Hash {
	return obj.hash
}

// IsSuccess returns true if there is a success, false otherwise
func (obj *expectation) IsSuccess() bool {
	return obj.success != nil
}

// Success returns the success, if any
func (obj *expectation) Success() outputs.Output {
	return obj.success
}

// IsMistake returns true if there is a mistake, false otherwise
func (obj *expectation) IsMistake() bool {
	return obj.mistake != nil
}

// Mistake returns the mistake, if any
func (obj *expectation) Mistake() links.Condition {
	return obj.mistake
}
