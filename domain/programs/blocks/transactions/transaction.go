package transactions

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions"
)

type transaction struct {
	hash       hash.Hash
	executions executions.Executions
	signature  signers.Signature
}

func createTransaction(
	hash hash.Hash,
	executions executions.Executions,
	signature signers.Signature,
) Transaction {
	out := transaction{
		hash:       hash,
		executions: executions,
		signature:  signature,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Executions returns the executions
func (obj *transaction) Executions() executions.Executions {
	return obj.executions
}

// Signature returns the signature
func (obj *transaction) Signature() signers.Signature {
	return obj.signature
}
