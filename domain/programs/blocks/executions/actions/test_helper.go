package actions

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/executions/actions/resources"
)

// NewActionsForTests creates actions for tests
func NewActionsForTests(list []Action) Actions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewActionWithCreateForTests creates an action with create for tests
func NewActionWithCreateForTests(create resources.Resource) Action {
	ins, err := NewActionBuilder().Create().WithCreate(create).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewActionWithDeleteForTests creates an action with delete for tests
func NewActionWithDeleteForTests(del hash.Hash) Action {
	ins, err := NewActionBuilder().Create().WithDelete(del).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
