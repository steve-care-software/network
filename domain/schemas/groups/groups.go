package groups

import (
	"errors"
	"fmt"
)

type groups struct {
	mp   map[string]Group
	list []Group
}

func createGroups(
	mp map[string]Group,
	list []Group,
) Groups {
	out := groups{
		mp:   mp,
		list: list,
	}

	return &out
}

// List returns the list
func (obj *groups) List() []Group {
	return obj.list
}

// Fetch fetches a group by name
func (obj *groups) Fetch(name string) (Group, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the groups does not contain a group of that name (%s)", name)
	return nil, errors.New(str)
}
