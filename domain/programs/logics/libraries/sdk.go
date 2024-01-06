package libraries

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the library builder
type Builder interface {
	Create() Builder
	WithLayers(layers layers.Layers) Builder
	WithLinks(links links.Links) Builder
	Now() (Library, error)
}

// Library represents the library
type Library interface {
	Hash() hash.Hash
	Layers() layers.Layers
	HasLinks() bool
	Links() links.Links
}
