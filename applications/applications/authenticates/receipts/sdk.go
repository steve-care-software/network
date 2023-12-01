package receipts

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the receipt application
type Application interface {
	Amount() (uint, error)
	List(index uint, amount uint) ([]hash.Hash, error)
	ListBySigner(pubKey signers.PublicKey, index uint, amount uint) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (receipts.Receipt, error)
	Insert(receipt receipts.Receipt) error
	Delete(hash hash.Hash) error
}
