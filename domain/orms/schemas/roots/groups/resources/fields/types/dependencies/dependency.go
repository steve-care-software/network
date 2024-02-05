package dependencies

type dependency struct {
	retriever string
	groups    []string
	resource  string
	kind      uint8
}

func createDependency(
	retriever string,
	groups []string,
	resource string,
	kind uint8,
) Dependency {
	out := dependency{
		retriever: retriever,
		groups:    groups,
		resource:  resource,
		kind:      kind,
	}

	return &out
}

// Retriever returns the retriever
func (obj *dependency) Retriever() string {
	return obj.retriever
}

// Groups returns the groups
func (obj *dependency) Groups() []string {
	return obj.groups
}

// Resource returns the resource
func (obj *dependency) Resource() string {
	return obj.resource
}

// Kind returns the kind
func (obj *dependency) Kind() uint8 {
	return obj.kind
}
