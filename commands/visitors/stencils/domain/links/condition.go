package links

import "steve.care/network/libraries/hash"

type condition struct {
	hash     hash.Hash
	resource ConditionResource
	operator Operator
	next     ConditionValue
}

func createCondition(
	hash hash.Hash,
	resource ConditionResource,
	operator Operator,
	next ConditionValue,
) Condition {
	out := condition{
		hash:     hash,
		resource: resource,
		operator: operator,
		next:     next,
	}

	return &out
}

// Hash returns the hash
func (obj *condition) Hash() hash.Hash {
	return obj.hash
}

// Resource returns the resource
func (obj *condition) Resource() ConditionResource {
	return obj.resource
}

// Operator returns the operator
func (obj *condition) Operator() Operator {
	return obj.operator
}

// Next returns the next value
func (obj *condition) Next() ConditionValue {
	return obj.next
}
