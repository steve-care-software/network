package expectations

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/databases/transactions"
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
	WithTransaction(trx transactions.Transaction) RepositoryBuilder
	WithQuery(query queries.Query) RepositoryBuilder
	WithCredentials(credentials credentials.Credentials) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a suite repository
type Repository interface {
	Amount() (uint, error)
	List() ([]hash.Hash, error)
	RetrieveByHash(hash hash.Hash) (Expectation, error)
	RetrieveByOutputLayer(layer hash.Hash) (Expectation, error)
	RetrieveByCondition(condition hash.Hash) (Expectation, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithTransaction(trx transactions.Transaction) ServiceBuilder
	WithQuery(query queries.Query) ServiceBuilder
	WithCredentials(credentials credentials.Credentials) ServiceBuilder
	Now() (Service, error)
}

// Service represents a suite service
type Service interface {
	Insert(ins Expectation) error
	Delete(hash hash.Hash) error
}
