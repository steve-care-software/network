package roots

import (
	"steve.care/network/domain/orms/schemas/roots/groups"
	"steve.care/network/domain/orms/schemas/roots/groups/resources"
	"steve.care/network/domain/orms/schemas/roots/methods"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewRootBuilder creates a new root builder
func NewRootBuilder() RootBuilder {
	return createRootBuilder()
}

// Builder represents a roots builder
type Builder interface {
	Create() Builder
	WithList(list []Root) Builder
	Now() (Roots, error)
}

// Roots represents roots
type Roots interface {
	List() []Root
	Search(path []string) (resources.Resource, error)
}

// RootBuilder represents the root builder
type RootBuilder interface {
	Create() RootBuilder
	WithName(name string) RootBuilder
	WithChains(chains groups.MethodChains) RootBuilder
	WithMethods(methods methods.Methods) RootBuilder
	Now() (Root, error)
}

// Root represents the root
type Root interface {
	Name() string
	Chains() groups.MethodChains
	Methods() methods.Methods
	Search(path []string) (resources.Resource, error)
}
