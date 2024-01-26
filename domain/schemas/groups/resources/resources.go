package resources

import (
	"errors"
	"fmt"
)

type resources struct {
	mp   map[string]Resource
	list []Resource
}

func createResources(
	mp map[string]Resource,
	list []Resource,
) Resources {
	out := resources{
		mp:   mp,
		list: list,
	}

	return &out
}

// List returns the list
func (obj *resources) List() []Resource {
	return obj.list
}

// Fetch fetches a resource by name
func (obj *resources) Fetch(name string) (Resource, error) {
	fmt.Printf("\n%v\n%s\n", obj.mp, name)
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the resources does not contain a resource of that name (%s)", name)
	return nil, errors.New(str)
}
