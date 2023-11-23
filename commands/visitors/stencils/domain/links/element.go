package links

import "steve.care/network/libraries/hash"

type element struct {
	hash      hash.Hash
	container []string
	condition Condition
}

func createElement(
	hash hash.Hash,
	container []string,
) Element {
	return createElementInternally(hash, container, nil)
}

func createElementWithCondition(
	hash hash.Hash,
	container []string,
	condition Condition,
) Element {
	return createElementInternally(hash, container, condition)
}

func createElementInternally(
	hash hash.Hash,
	container []string,
	condition Condition,
) Element {
	out := element{
		hash:      hash,
		container: container,
		condition: condition,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// Container returns the container
func (obj *element) Container() []string {
	return obj.container
}

// HasCondition returns true if there is a condition, false otheriwse
func (obj *element) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *element) Condition() Condition {
	return obj.condition
}
