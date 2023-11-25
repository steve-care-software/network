package layers

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the layer application
type Application interface {
	List() ([]hash.Hash, error)
	Exists(hash hash.Hash) (bool, error)
	Retrieve(hash hash.Hash) (layers.Layer, error)
	Insert(layer layers.Layer) error
	Delete(hash hash.Hash) error
}
