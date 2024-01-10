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

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := builder{
		hashAdapter:  hashAdapter,
		transactions: nil,
		parent:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() ContentBuilder {
	return createContentBuilder(
		app.hashAdapter,
	)
}

// WithTransactions add transactions to the builder
func (app *builder) WithTransactions(transactions transactions.Transactions) ContentBuilder {
	app.transactions = transactions
	return app
}

// WithParent add parent to the builder
func (app *builder) WithParent(parent hash.Hash) ContentBuilder {
	app.parent = parent
	return app
}

// Now builds a new Content instance
func (app *builder) Now() (Content, error) {
	if app.transactions != nil {
		return nil, errors.New("the transactions is mandatory in order to build a Content instance")
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
		return createContentWithParent(*pHash, app.transactions, app.parent), nil
	}

	return createContent(*pHash, app.transactions), nil
}
