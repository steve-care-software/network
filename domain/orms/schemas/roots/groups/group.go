package groups

import (
	"errors"
	"fmt"

	"steve.care/network/domain/orms/schemas/roots/groups/methods"
	"steve.care/network/domain/orms/schemas/roots/groups/resources"
)

type group struct {
	name    string
	chains  MethodChains
	methods methods.Methods
}

func createGroup(
	name string,
	chains MethodChains,
	methods methods.Methods,
) Group {
	out := group{
		name:    name,
		chains:  chains,
		methods: methods,
	}

	return &out
}

// Name returns the name
func (obj *group) Name() string {
	return obj.name
}

// Chains returns the chains
func (obj *group) Chains() MethodChains {
	return obj.chains
}

// Methods returns the methods
func (obj *group) Methods() methods.Methods {
	return obj.methods
}

// Search searches a resource by path
func (obj *group) Search(path []string) (resources.Resource, error) {
	if len(path) <= 0 {
		str := fmt.Sprintf("the path cannot be empty in order to search for a Resource in a Group (name: %s)", obj.name)
		return nil, errors.New(str)
	}

	if obj.name != path[0] {
		str := fmt.Sprintf("the first path element (%s) does not math the group name (%s)", path[0], obj.name)
		return nil, errors.New(str)
	}

	return obj.chains.Search(path[1:])
}
