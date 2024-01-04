package libraries

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
)

// Library represents the library
type Library interface {
	Hash() hash.Hash
	HasLayers() bool
	Layers() layers.Layers
	HasLinks() bool
	Links() links.Links
}
