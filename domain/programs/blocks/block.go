package blocks

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions"
)

type block struct {
	hash         hash.Hash
	transactions transactions.Transactions
	parent       hash.Hash
}

func createBlock(
	hash hash.Hash,
	transactions transactions.Transactions,
) Block {
	return createBlockInternally(hash, transactions, nil)
}

func createBlockWithParent(
	hash hash.Hash,
	transactions transactions.Transactions,
	parent hash.Hash,
) Block {
	return createBlockInternally(hash, transactions, parent)
}

func createBlockInternally(
	hash hash.Hash,
	transactions transactions.Transactions,
	parent hash.Hash,
) Block {
	out := block{
		hash:         hash,
		transactions: transactions,
		parent:       parent,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.hash
}

// Transactions returns the transactions
func (obj *block) Transactions() transactions.Transactions {
	return obj.transactions
}

// HasParent returns true if there is parent, false otherwise
func (obj *block) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *block) Parent() hash.Hash {
	return obj.parent
}
