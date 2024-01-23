package blocks

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions"
)

// NewBlockForTests creates a new block for tests
func NewBlockForTests(content Content, result []byte) Block {
	ins, err := NewBuilder().Create().WithContent(content).WithResult(result).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithParentForTests creates new content with parent for tests
func NewContentWithParentForTests(trx transactions.Transactions, parent hash.Hash) Content {
	ins, err := NewContentBuilder().Create().WithTransactions(trx).WithParent(parent).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentForTests creates new content for tests
func NewContentForTests(trx transactions.Transactions) Content {
	ins, err := NewContentBuilder().Create().WithTransactions(trx).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
