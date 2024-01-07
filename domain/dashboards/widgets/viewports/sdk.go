package viewports

import (
	"steve.care/network/domain/dashboards/widgets/viewports/dimensions"
	"steve.care/network/domain/dashboards/widgets/viewports/positions"
	"steve.care/network/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the viewport builder
type Builder interface {
	Create() Builder
	WithLayer(layer uint) Builder
	WithPosition(position positions.Position) Builder
	WithDimension(dimension dimensions.Dimension) Builder
	Now() (Viewport, error)
}

// Viewport represents a viewport
type Viewport interface {
	Hash() hash.Hash
	Layer() uint
	Position() positions.Position
	Dimension() dimensions.Dimension
}
