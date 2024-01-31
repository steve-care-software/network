package groups

import (
	"steve.care/network/domain/schemas/roots/groups/methods"
	"steve.care/network/domain/schemas/roots/groups/resources"
)

// NewBuilder creates a new group builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewMethodChainsBuilder creates a new method chains builder
func NewMethodChainsBuilder() MethodChainsBuilder {
	return createMethodChainsBuilder()
}

// NewMethodChainBuilder creates a new method chain builder
func NewMethodChainBuilder() MethodChainBuilder {
	return createMethodChainBuilder()
}

// Builder represents a group builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithChains(chains MethodChains) Builder
	WithMethods(methods methods.Methods) Builder
	Now() (Group, error)
}

// Group represents a resource group
type Group interface {
	Name() string
	Chains() MethodChains
	Methods() methods.Methods
	Search(path []string) (resources.Resource, error)
}

// MethodChainsBuilder represents method chains builder
type MethodChainsBuilder interface {
	Create() MethodChainsBuilder
	WithList(list []MethodChain) MethodChainsBuilder
	Now() (MethodChains, error)
}

// MethodChains returns method chains
type MethodChains interface {
	List() []MethodChain
	Search(path []string) (resources.Resource, error)
}

// MethodChainBuilder represents a method chain builder
type MethodChainBuilder interface {
	Create() MethodChainBuilder
	WithCondition(condition string) MethodChainBuilder
	WithRetriever(retriever []string) MethodChainBuilder
	WithElement(element Element) MethodChainBuilder
	Now() (MethodChain, error)
}

// MethodChain represents a method chain
type MethodChain interface {
	Condition() string
	Retriever() []string
	Element() Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithGroup(group Group) ElementBuilder
	WithResource(resource resources.Resource) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	IsGroup() bool
	Group() Group
	IsResource() bool
	Resource() resources.Resource
}
