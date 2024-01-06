package executions

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions"
)

// Builder represents executions builder
type Builder interface {
	Create() Builder
	WithList(list []Execution) Builder
	Now() (Executions, error)
}

// Executions represents executions
type Executions interface {
	Hash() hash.Hash
	List() []Execution
}

// ExecutionBuilder represents the execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithActions(actions actions.Actions) ExecutionBuilder
	WithReceipt(receipt hash.Hash) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Hash() hash.Hash
	Actions() actions.Actions
	HasReceipt() bool
	Receipt() hash.Hash
}
