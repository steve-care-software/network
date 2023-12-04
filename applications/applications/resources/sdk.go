package resources

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases/criterias/conditions"
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/databases/transactions"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/resources"
)

// Builder represents the resources application builder
type Builder interface {
	Create() Builder
	WithQuery(query queries.Query) Builder
	WithTransaction(trx transactions.Transaction) Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the resources application
type Application interface {
	// AmountOfResourcesInEntity returns the amount of resources in entity
	AmountOfResourcesInEntity(entity string) (uint, error)

	// ListResourceHashesInEntity lists resource hashes in entity
	ListResourceHashesInEntity(entity string, index uint, amount uint) ([]hash.Hash, error)

	// ListResourceHashesInEntityByCondition lists resource hashes in entity by condition
	ListResourceHashesInEntityByCondition(entity string, condition conditions.Condition) ([]hash.Hash, error)

	// RetrieveResourceByCondition retrieves a resource by condition
	RetrieveResourceByCondition(entity string, condition conditions.Condition) (resources.Resource, error)

	// InsertResource inserts a resource
	InsertResource(ins resources.Resource) error

	// DeleteResourceByEntityAndHash deletes a resource by entity and hash
	DeleteResourceByEntityAndHash(entity string, hash hash.Hash) error
}
