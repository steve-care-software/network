package groups

import "steve.care/network/domain/schemas/groups/resources"

type element struct {
	groups    Groups
	resources resources.Resources
}

func createElementWithResources(
	resources resources.Resources,
) Element {
	return createElementInternally(nil, resources)
}

func createElementWithGroups(
	groups Groups,
) Element {
	return createElementInternally(groups, nil)
}

func createElementInternally(
	groups Groups,
	resources resources.Resources,
) Element {
	out := element{
		groups:    groups,
		resources: resources,
	}

	return &out
}

// IsGroups returns true if there is groups, false otherwise
func (obj *element) IsGroups() bool {
	return obj.groups != nil
}

// Groups returns the groups, if any
func (obj *element) Groups() Groups {
	return obj.groups
}

// IsResources returns true if there is resources, false otherwise
func (obj *element) IsResources() bool {
	return obj.resources != nil
}

// Resources returns the resources, if any
func (obj *element) Resources() resources.Resources {
	return obj.resources
}
