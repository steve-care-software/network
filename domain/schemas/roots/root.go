package roots

import (
	"steve.care/network/domain/schemas/roots/groups"
	"steve.care/network/domain/schemas/roots/methods"
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
