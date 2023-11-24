package commands

import (
	"steve.care/network/domain/hash"
)

type commands struct {
	hash hash.Hash
	list []Command
}

func createCommands(
	hash hash.Hash,
	list []Command,
) Commands {
	out := commands{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *commands) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *commands) List() []Command {
	return obj.list
}
