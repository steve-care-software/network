package suites

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/suites"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the link application
type Application interface {
	Amount() (uint, error)
	List(index uint, amount uint) ([]hash.Hash, error)
	Exists(hash hash.Hash) (bool, error)
	RetrieveByHash(hash hash.Hash) (suites.Suite, error)
	RetrieveByOrigin(executedLayers []hash.Hash) (suites.Suite, error)
	RetrieveByExpectation(hash hash.Hash) (suites.Suite, error)
	Insert(suite suites.Suite) error
	Delete(hash hash.Hash) error
}
