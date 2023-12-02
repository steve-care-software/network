package expectations

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers"
	"steve.care/network/domain/receipts/commands/links"
)

// Builder represents an expectation builder
type Builder interface {
	Create() Builder
	WithOutput(output layers.Layer) Builder
	WithCondition(condition links.Condition) Builder
	Now() (Expectation, error)
}

// Expectation represents an expectation
type Expectation interface {
	Hash() hash.Hash
	IsOutput() bool
	Output() layers.Layer
	IsCondition() bool
	Condition() links.Condition
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
	Retrieve(hash hash.Hash) (Expectation, error)
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
	Insert(ins Expectation) error
	Delete(hash hash.Hash) error
}
