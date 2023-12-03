package resources

import "steve.care/network/domain/databases/criterias/conditions"

type resource struct {
	container string
	condition conditions.Condition
}

func createResource(
	container string,
	condition conditions.Condition,
) Resource {
	out := resource{
		container: container,
		condition: condition,
	}

	return &out
}

// Container returns the container
func (obj *resource) Container() string {
	return obj.container
}

// Condition returns the condition
func (obj *resource) Condition() conditions.Condition {
	return obj.condition
}
