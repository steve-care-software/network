package viewports

import (
	"steve.care/network/domain/dashboards/widgets/viewports/dimensions"
	"steve.care/network/domain/dashboards/widgets/viewports/positions"
	"steve.care/network/domain/hash"
)

// Builder represents the viewport builder
type Builder interface {
	Create() Builder
	WithPosition(position positions.Position) Builder
	WithDimension(dimension dimensions.Dimension) Builder
	Now() (Viewport, error)
}

// Viewport represents a viewport
type Viewport interface {
	Hash() hash.Hash
	Position() positions.Position
	Dimension() dimensions.Dimension
}
