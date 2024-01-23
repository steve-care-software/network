package transactions

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/programs/blocks/transactions/executions"
)

// NewTransactionsForTests creates a new transactions for tests
func NewTransactionsForTests(list []Transaction) Transactions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTransactionForTests creates a new transaction for tests
func NewTransactionForTests(executions executions.Executions, signature signers.Signature) Transaction {
	ins, err := NewTransactionBuilder().Create().WithExecutions(executions).WithSignature(signature).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
