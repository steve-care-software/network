package executions

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions"
)

// Executions represents executions
type Executions interface {
	Hash() hash.Hash
	List() []Execution
}

// Execution represents an execution
type Execution interface {
	Hash() hash.Hash
	Actions() actions.Actions
	Receipt() hash.Hash
}
