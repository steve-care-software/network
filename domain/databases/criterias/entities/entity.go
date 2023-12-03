package entities

import "steve.care/network/domain/databases/criterias/entities/resources"

type entity struct {
	resource resources.Resource
	fields   []string
}

func createEntity(
	resource resources.Resource,
	fields []string,
) Entity {
	out := entity{
		resource: resource,
		fields:   fields,
	}

	return &out
}

// Resource returns the resource
func (obj *entity) Resource() resources.Resource {
	return obj.resource
}

// Fields returns the fields
func (obj *entity) Fields() []string {
	return obj.fields
}
