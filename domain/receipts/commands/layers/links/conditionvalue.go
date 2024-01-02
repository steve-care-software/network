package links

import "steve.care/network/domain/hash"

type conditionValue struct {
	hash      hash.Hash
	resource  ConditionResource
	condition Condition
}

func createConditionValueWithResource(
	hash hash.Hash,
	resource ConditionResource,
) ConditionValue {
	return createConditionValueInternally(hash, resource, nil)
}

func createConditionValueWithCondition(
	hash hash.Hash,
	condition Condition,
) ConditionValue {
	return createConditionValueInternally(hash, nil, condition)
}

func createConditionValueInternally(
	hash hash.Hash,
	resource ConditionResource,
	condition Condition,
) ConditionValue {
	out := conditionValue{
		hash:      hash,
		resource:  resource,
		condition: condition,
	}

	return &out
}

// Hash returns the hash
func (obj *conditionValue) Hash() hash.Hash {
	return obj.hash
}

// IsResource returns true if there is a resource, false otherwise
func (obj *conditionValue) IsResource() bool {
	return obj.resource != nil
}

// Resource returns the resource, if any
func (obj *conditionValue) Resource() ConditionResource {
	return obj.resource
}

// IsCondition returns true if there is an condition, false otherwise
func (obj *conditionValue) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *conditionValue) Condition() Condition {
	return obj.condition
}
