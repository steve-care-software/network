package resources

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/databases/resources/tokens"
	"steve.care/network/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the resource builder
type Builder interface {
	Create() Builder
	WithToken(token tokens.Token) Builder
	WithSignature(signature signers.Signature) Builder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	Token() tokens.Token
	Signature() signers.Signature
}

// Repository represents a resource repository
type Repository interface {
	// Amount returns the amount of resources
	Amount() (uint, error)

	// AmountByCriteria returns the amount of resources by criteria
	AmountByCriteria(criteria hash.Hash) (uint, error)

	// ListByCriteria lists resource hashes by criteria
	ListByCriteria(criteria hash.Hash) ([]hash.Hash, error)

	// RetrieveByCriteria retrieves a resource by criteria
	RetrieveByCriteria(criteria hash.Hash) (Resource, error)
}

// Service represents a resource service
type Service interface {
	// Insert inserts a resource
	Insert(ins Resource) error

	// Delete deletes a resource
	Delete(hash hash.Hash) error
}
