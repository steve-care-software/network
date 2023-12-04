package resources

import (
	"time"

	"steve.care/network/domain/databases/criterias/conditions"
	"steve.care/network/domain/databases/criterias/entities/resources"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
	"steve.care/network/domain/resources/layers"
	"steve.care/network/domain/resources/links"
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
	IsLayer() bool
	Layer() layers.Layer
	IsLink() bool
	Link() links.Link
	IsSuite() bool
	Suite() suites.Suite
	IsReceipt() bool
	Receipt() receipts.Receipt
}

// Repository represents a resource repository
type Repository interface {
	// AmountOfResourcesInEntity returns the amount of resources in entity
	AmountOfResourcesInEntity(entity string) (uint, error)

	// ListResourceHashesInEntity lists resource hashes in entity
	ListResourceHashesInEntity(entity string, index uint, amount uint) ([]hash.Hash, error)

	// ListResourceHashesInEntityByCondition lists resource hashes in entity by condition
	ListResourceHashesInEntityByCondition(entity string, condition conditions.Condition) ([]hash.Hash, error)

	// RetrieveResourceByCondition retrieves a resource by condition
	RetrieveResourceByCondition(entity string, condition conditions.Condition) (resources.Resource, error)
}

// Service represents a resource service
type Service interface {
	// InsertResource inserts a resource
	InsertResource(ins resources.Resource) error

	// DeleteResourceByEntityAndHash deletes a resource by entity and hash
	DeleteResourceByEntityAndHash(entity string, hash hash.Hash) error
}
