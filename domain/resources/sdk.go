package resources

import (
	"time"

	"steve.care/network/domain/databases/criterias/conditions"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/resources/layers"
	"steve.care/network/domain/resources/links"
	"steve.care/network/domain/resources/queries"
	"steve.care/network/domain/resources/suites"
)

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	CreatedBy() receipts.Receipt
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

	// AmountInEntity returns the amount of resources in entity
	AmountInEntity(entity string) (uint, error)

	// List lists resources
	List(index uint, amount uint) []hash.Hash

	// ListInEntity lists resource hashes in entity
	ListInEntity(entity string, index uint, amount uint) ([]hash.Hash, error)

	// ListInEntityByCondition lists resource hashes in entity by condition
	ListInEntityByCondition(entity string, condition conditions.Condition) ([]hash.Hash, error)

	// RetrieveByCondition retrieves a resource by condition
	RetrieveByCondition(entity string, condition conditions.Condition) (Resource, error)
}

// Service represents a resource service
type Service interface {
	// Insert inserts a resource
	Insert(ins Resource) error

	// Delete deletes a resource
	Delete(hash hash.Hash) error
}
