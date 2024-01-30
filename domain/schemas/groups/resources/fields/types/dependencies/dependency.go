package dependencies

type dependency struct {
	groups   []string
	resource string
}

func createDependency(
	groups []string,
	resource string,
) Dependency {
	out := dependency{
		groups:   groups,
		resource: resource,
	}

	return &out
}

// Groups returns the groups
func (obj *dependency) Groups() []string {
	return obj.groups
}

// Resource returns the resource
func (obj *dependency) Resource() string {
	return obj.resource
}
