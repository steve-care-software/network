package transactions

import "steve.care/network/domain/hash"

type transactions struct {
	hash hash.Hash
	list []Transaction
}

func createTransactions(
	hash hash.Hash,
	list []Transaction,
) Transactions {
	out := transactions{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *transactions) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *transactions) List() []Transaction {
	return obj.list
}
