package resources

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases/criterias/conditions"
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/databases/resources"
	"steve.care/network/domain/databases/transactions"
	"steve.care/network/domain/hash"
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
	RetrieveByCondition(entity string, condition conditions.Condition) (resources.Resource, error)

	// Insert inserts a resource
	Insert(ins resources.Resource) error

	// Delete deletes a resource
	Delete(hash hash.Hash) error
}
