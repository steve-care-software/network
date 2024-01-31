package groups

import "steve.care/network/domain/schemas/roots/groups/resources"

type element struct {
	group    Group
	resource resources.Resource
}

func createElementWithResource(
	resource resources.Resource,
) Element {
	return createElementInternally(nil, resource)
}

func createElementWithGroup(
	group Group,
) Element {
	return createElementInternally(group, nil)
}

func createElementInternally(
	group Group,
	resource resources.Resource,
) Element {
	out := element{
		group:    group,
		resource: resource,
	}

	return &out
}

// IsGroup returns true if there is group, false otherwise
func (obj *element) IsGroup() bool {
	return obj.group != nil
}

// Group returns the group, if any
func (obj *element) Group() Group {
	return obj.group
}

// IsResource returns true if there is resource, false otherwise
func (obj *element) IsResource() bool {
	return obj.resource != nil
}

// Resource returns the resource, if any
func (obj *element) Resource() resources.Resource {
	return obj.resource
}
