package transactions

import (
	"errors"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions"
)

type transactionBuilder struct {
	hashAdapter hash.Adapter
	executions  executions.Executions
	signature   signers.Signature
}

func createTransactionBuilder(
	hashAdapter hash.Adapter,
) TransactionBuilder {
	out := transactionBuilder{
		hashAdapter: hashAdapter,
		executions:  nil,
		signature:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *transactionBuilder) Create() TransactionBuilder {
	return createTransactionBuilder(
		app.hashAdapter,
	)
}

// WithExecutions add executions to the builder
func (app *transactionBuilder) WithExecutions(executions executions.Executions) TransactionBuilder {
	app.executions = executions
	return app
}

// WithSignature add signature to the builder
func (app *transactionBuilder) WithSignature(signature signers.Signature) TransactionBuilder {
	app.signature = signature
	return app
}

// Now builds a new Transaction instance
func (app *transactionBuilder) Now() (Transaction, error) {
	if app.executions == nil {
		return nil, errors.New("the executions is mandatory in order to build a Transaction instance")
	}

	if app.signature == nil {
		return nil, errors.New("the signature is mandatory in order to build a Transaction instance")
	}

	sigBytes, err := app.signature.Bytes()
	if err != nil {
		return nil, err
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.executions.Hash().Bytes(),
		sigBytes,
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(*pHash, app.executions, app.signature), nil
}
