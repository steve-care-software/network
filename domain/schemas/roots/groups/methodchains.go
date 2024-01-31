package groups

import (
	"errors"
	"fmt"
	"strings"

	"steve.care/network/domain/schemas/roots/groups/resources"
)

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

// Search searches a resource by path
func (obj *methodChains) Search(path []string) (resources.Resource, error) {
	for _, oneMethodChain := range obj.list {
		element := oneMethodChain.Element()
		if element.IsGroup() {
			group := element.Group()
			return group.Search(path)
		}

		if element.IsResource() {
			if len(path) != 1 {
				continue
			}

			name := element.Resource().Name()
			if name != path[0] {
				continue
			}

			return element.Resource(), nil
		}
	}

	str := fmt.Sprintf("there is no Resource discovered at the provided path: %s", strings.Join(path, "/"))
	return nil, errors.New(str)
}
