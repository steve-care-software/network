package groups

type groups struct {
	list []Group
}

func createGroups(
	list []Group,
) Groups {
	out := groups{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *groups) List() []Group {
	return obj.list
}
