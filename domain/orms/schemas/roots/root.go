package roots

import (
	"errors"
	"fmt"

	"steve.care/network/domain/orms/schemas/roots/groups"
	"steve.care/network/domain/orms/schemas/roots/groups/resources"
	"steve.care/network/domain/orms/schemas/roots/methods"
)

type root struct {
	name    string
	chains  groups.MethodChains
	methods methods.Methods
}

func createRoot(
	name string,
	chains groups.MethodChains,
	methods methods.Methods,
) Root {
	out := root{
		name:    name,
		chains:  chains,
		methods: methods,
	}

	return &out
}

// Name returns the name
func (obj *root) Name() string {
	return obj.name
}

// Chains returns the chains
func (obj *root) Chains() groups.MethodChains {
	return obj.chains
}

// Methods returns the methods
func (obj *root) Methods() methods.Methods {
	return obj.methods
}

// Search searches a resource by path
func (obj *root) Search(path []string) (resources.Resource, error) {
	if len(path) <= 0 {
		str := fmt.Sprintf("the path cannot be empty in order to search for a Resource in a Root (name: %s)", obj.name)
		return nil, errors.New(str)
	}

	if obj.name != path[0] {
		str := fmt.Sprintf("the first path element (%s) does not math the root name (%s)", path[0], obj.name)
		return nil, errors.New(str)
	}

	return obj.chains.Search(path[1:])
}
