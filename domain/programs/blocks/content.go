package blocks

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions"
)

type content struct {
	hash         hash.Hash
	transactions transactions.Transactions
	parent       hash.Hash
}

func createContent(
	hash hash.Hash,
	transactions transactions.Transactions,
) Content {
	return createContentInternally(hash, transactions, nil)
}

func createContentWithParent(
	hash hash.Hash,
	transactions transactions.Transactions,
	parent hash.Hash,
) Content {
	return createContentInternally(hash, transactions, parent)
}

func createContentInternally(
	hash hash.Hash,
	transactions transactions.Transactions,
	parent hash.Hash,
) Content {
	out := content{
		hash:         hash,
		transactions: transactions,
		parent:       parent,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Transactions returns the transactions
func (obj *content) Transactions() transactions.Transactions {
	return obj.transactions
}

// HasParent returns true if there is parent, false otherwise
func (obj *content) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *content) Parent() hash.Hash {
	return obj.parent
}
