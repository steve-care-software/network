package widgets

import (
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs"
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
	Viewport() viewports.Viewport
	Program() programs.Program
	Input() []byte
}
