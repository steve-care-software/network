package roots

import (
	"steve.care/network/domain/schemas/roots/groups"
	"steve.care/network/domain/schemas/roots/groups/resources"
	"steve.care/network/domain/schemas/roots/methods"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the root builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithChains(chains groups.MethodChains) Builder
	WithMethods(methods methods.Methods) Builder
	Now() (Root, error)
}

// Root represents the root
type Root interface {
	Name() string
	Chains() groups.MethodChains
	Methods() methods.Methods
	Search(path []string) (resources.Resource, error)
}
