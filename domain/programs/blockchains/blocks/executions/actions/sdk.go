package actions

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources"
)

// Actions represents the actions
type Actions interface {
	Hash() hash.Hash
	List() []Action
}

// Action represents an execution action
type Action interface {
	Hash() hash.Hash
	IsCreate() bool
	Create() resources.Resource
	IsDelete() bool
	Delete() hash.Hash
}
