package expectations

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
)

type expectation struct {
	hash      hash.Hash
	output    layers.Layer
	condition links.Condition
}

func createExpectationWithOutput(
	hash hash.Hash,
	output layers.Layer,
) Expectation {
	return createExpectationInternally(
		hash,
		output,
		nil,
	)
}

func createExpectationWithCondition(
	hash hash.Hash,
	condition links.Condition,
) Expectation {
	return createExpectationInternally(
		hash,
		nil,
		condition,
	)
}

func createExpectationInternally(
	hash hash.Hash,
	output layers.Layer,
	condition links.Condition,
) Expectation {
	out := expectation{
		hash:      hash,
		output:    output,
		condition: condition,
	}

	return &out
}

// Hash returns the hash
func (obj *expectation) Hash() hash.Hash {
	return obj.hash
}

// IsOutput returns true if there is an output, false otherwise
func (obj *expectation) IsOutput() bool {
	return obj.output != nil
}

// Output returns the output, if any
func (obj *expectation) Output() layers.Layer {
	return obj.output
}

// IsCondition returns true if there is a condition, false otherwise
func (obj *expectation) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *expectation) Condition() links.Condition {
	return obj.condition
}
