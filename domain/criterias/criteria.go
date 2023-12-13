package criterias

import (
	"steve.care/network/domain/criterias/conditions"
	"steve.care/network/domain/hash"
)

type criteria struct {
	hash      hash.Hash
	entity    string
	condition conditions.Condition
}

func createCriteria(
	hash hash.Hash,
	entity string,
	condition conditions.Condition,
) Criteria {
	out := criteria{
		hash:      hash,
		entity:    entity,
		condition: condition,
	}

	return &out
}

// Hash returns the hash
func (obj *criteria) Hash() hash.Hash {
	return obj.hash
}

// Entity returns the entity
func (obj *criteria) Entity() string {
	return obj.entity
}

// Condition returns the condition
func (obj *criteria) Condition() conditions.Condition {
	return obj.condition
}
