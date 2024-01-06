package viewports

import (
	"steve.care/network/domain/dashboards/widgets/viewports/dimensions"
	"steve.care/network/domain/dashboards/widgets/viewports/positions"
	"steve.care/network/domain/hash"
)

// Viewport represents a viewport
type Viewport interface {
	Hash() hash.Hash
	Position() positions.Position
	Dimension() dimensions.Dimension
}
