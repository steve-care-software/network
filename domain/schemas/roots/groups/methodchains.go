package groups

type methodChains struct {
	list []MethodChain
}

func createMethodChains(
	list []MethodChain,
) MethodChains {
	out := methodChains{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *methodChains) List() []MethodChain {
	return obj.list
}
