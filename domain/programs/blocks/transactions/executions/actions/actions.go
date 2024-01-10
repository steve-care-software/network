package actions

import "steve.care/network/domain/hash"

type actions struct {
	hash hash.Hash
	list []Action
}

func createActions(
	hash hash.Hash,
	list []Action,
) Actions {
	out := actions{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *actions) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *actions) List() []Action {
	return obj.list
}
