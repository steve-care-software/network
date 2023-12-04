package resources

import (
	"steve.care/network/domain/credentials"
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

	// AmountByCriteria returns the amount of resources by criteria
	AmountByCriteria(criteria hash.Hash) (uint, error)

	// ListByCriteria lists resource hashes by criteria
	ListByCriteria(criteria hash.Hash) ([]hash.Hash, error)

	// RetrieveByCriteria retrieves a resource by criteria
	RetrieveByCriteria(criteria hash.Hash) (resources.Resource, error)

	// Insert inserts a resource
	Insert(ins resources.Resource) error

	// Delete deletes a resource
	Delete(hash hash.Hash) error
}