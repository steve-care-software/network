package resources

import "steve.care/network/domain/databases/criterias/conditions"

type resource struct {
	entity    string
	condition conditions.Condition
}

func createResource(
	entity string,
	condition conditions.Condition,
) Resource {
	out := resource{
		entity:    entity,
		condition: condition,
	}

	return &out
}

// Entity returns the entity
func (obj *resource) Entity() string {
	return obj.entity
}

// Condition returns the condition
func (obj *resource) Condition() conditions.Condition {
	return obj.condition
}
