package links

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/links"
)

// NewBuilder creates a new application builder
func NewBuilder(
	repositoryBuilder links.RepositoryBuilder,
	serviceBuilder links.ServiceBuilder,
) Builder {
	return createBuilder(
		repositoryBuilder,
		serviceBuilder,
	)
}

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
	RetrieveByHash(hash hash.Hash) (links.Link, error)
	RetrieveByOrigin(executedLayers []hash.Hash) (links.Link, error)
	Insert(link links.Link) error
	Delete(hash hash.Hash) error
}
