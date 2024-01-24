package groups

type group struct {
	name     string
	elements Elements
}

func createGroup(
	name string,
	elements Elements,
) Group {
	out := group{
		name:     name,
		elements: elements,
	}

	return &out
}

// Name returns the name
func (obj *group) Name() string {
	return obj.name
}

// Elements returns the elements
func (obj *group) Elements() Elements {
	return obj.elements
}
