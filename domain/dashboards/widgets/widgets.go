package widgets

import "steve.care/network/domain/hash"

type widgets struct {
	hash hash.Hash
	list []Widget
}

func createWidgets(
	hash hash.Hash,
	list []Widget,
) Widgets {
	out := widgets{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *widgets) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *widgets) List() []Widget {
	return obj.list
}
