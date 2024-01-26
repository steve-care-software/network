package groups

type group struct {
	name   string
	chains MethodChains
}

func createGroup(
	name string,
	chains MethodChains,
) Group {
	out := group{
		name:   name,
		chains: chains,
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
