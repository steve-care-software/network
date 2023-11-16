package commands

import "steve.care/network/libraries/hash"

// Builder represents a resources builder
type Builder interface {
	Create() Builder
	WithList(list []Command) Builder
	Now() (Commands, error)
}

// Commands represents the commands
type Commands interface {
	List() []Command
}

// CommandBuilder represents a command builder
type CommandBuilder interface {
	Create() CommandBuilder
	Now() (Command, error)
}

// Command represents a command
type Command interface {
	Hash() hash.Hash
}
