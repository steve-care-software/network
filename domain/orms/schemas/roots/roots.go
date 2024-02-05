package roots

import (
	"errors"
	"fmt"
	"strings"

	"steve.care/network/domain/orms/schemas/roots/groups/resources"
)

type roots struct {
	list []Root
}

func createRoots(
	list []Root,
) Roots {
	out := roots{
		list: list,
	}

	return &out
}

// List returns the roots
func (obj *roots) List() []Root {
	return obj.list
}

// Search searches a resource by path
func (obj *roots) Search(path []string) (resources.Resource, error) {
	for _, oneRoot := range obj.list {
		retResource, err := oneRoot.Search(path)
		if err != nil {
			return nil, err
		}

		return retResource, nil
	}

	str := fmt.Sprintf("there is no Resource related to the provided path: %s", strings.Join(path, "/"))
	return nil, errors.New(str)
}
