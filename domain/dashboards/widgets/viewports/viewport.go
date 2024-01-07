package viewports

import (
	"steve.care/network/domain/dashboards/widgets/viewports/dimensions"
	"steve.care/network/domain/dashboards/widgets/viewports/positions"
	"steve.care/network/domain/hash"
)

type viewport struct {
	hash      hash.Hash
	layer     uint
	position  positions.Position
	dimension dimensions.Dimension
}

func createViewport(
	hash hash.Hash,
	layer uint,
	position positions.Position,
	dimension dimensions.Dimension,
) Viewport {
	out := viewport{
		hash:      hash,
		layer:     layer,
		position:  position,
		dimension: dimension,
	}

	return &out
}

// Hash returns the hash
func (obj *viewport) Hash() hash.Hash {
	return obj.hash
}

// Layer returns the layer
func (obj *viewport) Layer() uint {
	return obj.layer
}

// Position returns the position
func (obj *viewport) Position() positions.Position {
	return obj.position
}

// Dimension returns the dimension
func (obj *viewport) Dimension() dimensions.Dimension {
	return obj.dimension
}
