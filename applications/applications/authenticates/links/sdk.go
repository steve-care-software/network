package links

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers"
	"steve.care/network/domain/receipts/commands/links"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() Builder
}

// Application represents the link application
type Application interface {
	List() ([]hash.Hash, error)
	RetrieveByHash(hash hash.Hash) (layers.Layer, error)
	RetrieveByOrigin(executedLayers []hash.Hash) (links.Link, error)
	Insert(link links.Link) error
	Delete(hash hash.Hash) error
}
