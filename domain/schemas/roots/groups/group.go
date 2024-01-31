package groups

import "steve.care/network/domain/schemas/roots/groups/methods"

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
