package executions

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions"
)

type execution struct {
	hash    hash.Hash
	actions actions.Actions
	receipt hash.Hash
}

func createExecution(
	hash hash.Hash,
	actions actions.Actions,
) Execution {
	return createExecutionInternally(hash, actions, nil)
}

func createExecutionWithReceipt(
	hash hash.Hash,
	actions actions.Actions,
	receipt hash.Hash,
) Execution {
	return createExecutionInternally(hash, actions, receipt)
}

func createExecutionInternally(
	hash hash.Hash,
	actions actions.Actions,
	receipt hash.Hash,
) Execution {
	out := execution{
		hash:    hash,
		actions: actions,
		receipt: receipt,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// Actions returns the actions
func (obj *execution) Actions() actions.Actions {
	return obj.actions
}

// HasReceipt returns true if there is a receipt, false otherwise
func (obj *execution) HasReceipt() bool {
	return obj.receipt != nil
}

// Receipt returns the receipt
func (obj *execution) Receipt() hash.Hash {
	return obj.receipt
}
