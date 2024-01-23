package executions

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions"
)

// NewExecutionsForTests creates a new executions for tests
func NewExecutionsForTests(list []Execution) Executions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithReceiptForTests creates a new execution with receipt for tests
func NewExecutionWithReceiptForTests(actions actions.Actions, receipt hash.Hash) Execution {
	ins, err := NewExecutionBuilder().Create().WithActions(actions).WithReceipt(receipt).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionForTests creates a new execution for tests
func NewExecutionForTests(actions actions.Actions) Execution {
	ins, err := NewExecutionBuilder().Create().WithActions(actions).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
