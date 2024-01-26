package groups

import "steve.care/network/domain/schemas/groups/resources"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewGroupBuilder creates a new group builder
func NewGroupBuilder() GroupBuilder {
	return createGroupBuilder()
}

// NewElementsBuilder creates a new elements builder
func NewElementsBuilder() ElementsBuilder {
	return createElementsBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// Builder represents the groups builder
type Builder interface {
	Create() Builder
	WithList(list []Group) Builder
	Now() (Groups, error)
}

// Groups represents groups
type Groups interface {
	List() []Group
	Fetch(name string) (Group, error)
}

// GroupBuilder represents a group builder
type GroupBuilder interface {
	Create() GroupBuilder
	WithName(name string) GroupBuilder
	WithElements(elements Elements) GroupBuilder
	Now() (Group, error)
}

// Group represents a resource group
type Group interface {
	Name() string
	Elements() Elements
}

// ElementsBuilder represents an elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithList(list []Element) ElementsBuilder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	List() []Element
	Search(name string) (Group, error)
	Resource(name string) (resources.Resource, error)
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithGroups(groups Groups) ElementBuilder
	WithResources(resources resources.Resources) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	IsGroups() bool
	Groups() Groups
	IsResources() bool
	Resources() resources.Resources
}
