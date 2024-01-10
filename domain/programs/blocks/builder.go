package blocks

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions"
)

type builder struct {
	hashAdapter  hash.Adapter
	transactions transactions.Transactions
	parent       hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		transactions: nil,
		parent:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithTransactions add transactions to the builder
func (app *builder) WithTransactions(transactions transactions.Transactions) Builder {
	app.transactions = transactions
	return app
}

// WithParent add parent to the builder
func (app *builder) WithParent(parent hash.Hash) Builder {
	app.parent = parent
	return app
}

// Now builds a new Block instance
func (app *builder) Now() (Block, error) {
	if app.transactions != nil {
		return nil, errors.New("the transactions is mandatory in order to build a Block instance")
	}

	data := [][]byte{
		app.transactions.Hash().Bytes(),
	}

	if app.parent != nil {
		data = append(data, app.parent.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		return createBlockWithParent(*pHash, app.transactions, app.parent), nil
	}

	return createBlock(*pHash, app.transactions), nil
}
