package actions

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources"
)

// Builder represents the actions builder
type Builder interface {
	Create() Builder
	WithList(list []Action) Builder
	Now() (Actions, error)
}

// Actions represents the actions
type Actions interface {
	Hash() hash.Hash
	List() []Action
}

// ActionBuilder represents the action builder
type ActionBuilder interface {
	Create() ActionBuilder
	WithCreate(create resources.Resource) ActionBuilder
	WithDelete(del hash.Hash) ActionBuilder
	Now() (Action, error)
}

// Action represents an execution action
type Action interface {
	Hash() hash.Hash
	IsCreate() bool
	Create() resources.Resource
	IsDelete() bool
	Delete() hash.Hash
}