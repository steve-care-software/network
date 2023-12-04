package resources

import (
	"time"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/databases/resources/layers"
	"steve.care/network/domain/databases/resources/links"
	"steve.care/network/domain/databases/resources/queries"
	"steve.care/network/domain/databases/resources/suites"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
)

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	Token() Token
	Signature() signers.Signature
}

// Token represents the token
type Token interface {
	Hash() hash.Hash
	Content() Content
	CreatedOn() time.Time
}

// Content represents a resource content
type Content interface {
	Hash() hash.Hash
	IsLayer() bool
	Layer() layers.Layer
	IsLink() bool
	Link() links.Link
	IsSuite() bool
	Suite() suites.Suite
	IsReceipt() bool
	Receipt() receipts.Receipt
	IsQuery() bool
	Query() queries.Query
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
