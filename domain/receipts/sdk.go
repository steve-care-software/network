package receipts

import (
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a receipt builder
type Builder interface {
	Create() Builder
	WithCommands(commands commands.Commands) Builder
	WithSignature(signature signers.Signature) Builder
	Now() (Receipt, error)
}

// Receipt represents a receipt
type Receipt interface {
	Hash() hash.Hash
	Commands() commands.Commands
	Signature() signers.Signature
}

// RepositoryBuilder represents the repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithAccount(account accounts.Account) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a receipt repository
type Repository interface {
	List(index uint, amount uint) ([]hash.Hash, error)
	ListBySigner(pubKey signers.PublicKey, index uint, amount uint) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Receipt, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithAccount(account accounts.Account) ServiceBuilder
	Now() (Service, error)
}

// Service represents a receipt service
type Service interface {
	Insert(receipt Receipt) error
	Delete(hash hash.Hash) error
}
