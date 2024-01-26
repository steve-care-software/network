package groups

import (
	"errors"
	"fmt"

	"steve.care/network/domain/schemas/groups/resources"
)

type elements struct {
	list []Element
}

func createElements(
	list []Element,
) Elements {
	out := elements{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *elements) List() []Element {
	return obj.list
}

// Search searches for a sub group
func (obj *elements) Search(name string) (Group, error) {
	for _, oneElement := range obj.list {
		if !oneElement.IsGroups() {
			continue
		}

		group, err := oneElement.Groups().Fetch(name)
		if err != nil {
			continue
		}

		return group, nil
	}

	str := fmt.Sprintf("there is no sub-group (name: %s) in the provided element", name)
	return nil, errors.New(str)
}

// Resource retrieves a resource by name
func (obj *elements) Resource(name string) (resources.Resource, error) {
	for _, oneElement := range obj.list {
		if !oneElement.IsResources() {
			continue
		}

		res, err := oneElement.Resources().Fetch(name)
		if err != nil {
			continue
		}

		return res, nil
	}

	str := fmt.Sprintf("there is no resource (name: %s) in the provided elements", name)
	return nil, errors.New(str)
}
