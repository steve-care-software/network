package criterias

import "steve.care/network/domain/databases/criterias/conditions"

type criteria struct {
	entity    string
	condition conditions.Condition
}

func createCriteria(
	entity string,
	condition conditions.Condition,
) Criteria {
	out := criteria{
		entity:    entity,
		condition: condition,
	}

	return &out
}

// Entity returns the entity
func (obj *criteria) Entity() string {
	return obj.entity
}

// Condition returns the condition
func (obj *criteria) Condition() conditions.Condition {
	return obj.condition
}
