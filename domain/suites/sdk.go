package suites

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers"
	"steve.care/network/domain/receipts/commands/links"
	"steve.care/network/domain/suites/expectations"
)

// Builder represents a suite builder
type Builder interface {
	Create() Builder
	WithOrigin(origin links.Origin) Builder
	WithInput(input layers.Layer) Builder
	WithExpectation(expectation expectations.Expectation) Builder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Hash() hash.Hash
	Origin() links.Origin
	Input() layers.Layer
	Expectation() expectations.Expectation
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithTransaction(trx databases.Transaction) RepositoryBuilder
	WithQuery(query databases.Query) RepositoryBuilder
	WithCredentials(credentials credentials.Credentials) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a suite repository
type Repository interface {
	Amount() (uint, error)
	List() ([]hash.Hash, error)
	RetrieveByHash(hash hash.Hash) (Suite, error)
	RetrieveByOrigin(origin hash.Hash) (Suite, error)
	RetrieveByInputLayer(layer hash.Hash) (Suite, error)
	RetrieveByExpectation(expectation hash.Hash) (Suite, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithTransaction(trx databases.Transaction) ServiceBuilder
	WithQuery(query databases.Query) ServiceBuilder
	WithCredentials(credentials credentials.Credentials) ServiceBuilder
	Now() (Service, error)
}

// Service represents a suite service
type Service interface {
	Insert(ins Suite) error
	Delete(hash hash.Hash) error
}
