package widgets

import (
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
)

// Widgets represents widgets
type Widgets interface {
	Hash() hash.Hash
	List() []Widget
}

// Widget represents a dashboard widget
type Widget interface {
	Hash() hash.Hash
	Title() string
	Program() hash.Hash
	Input() []byte
	HasViewport() bool
	Viewport() viewports.Viewport
}
