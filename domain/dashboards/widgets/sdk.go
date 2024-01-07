package widgets

import (
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewWidgetBuilder creates a new widget builder
func NewWidgetBuilder() WidgetBuilder {
	hashAdapter := hash.NewAdapter()
	return createWidgetBuilder(
		hashAdapter,
	)
}

// Builder represents the widgets builder
type Builder interface {
	Create() Builder
	WithList(list []Widget) Builder
	Now() (Widgets, error)
}

// Widgets represents widgets
type Widgets interface {
	Hash() hash.Hash
	List() []Widget
}

// WidgetBuilder represents the widget builder
type WidgetBuilder interface {
	Create() WidgetBuilder
	WithTitle(title string) WidgetBuilder
	WithProgram(program hash.Hash) WidgetBuilder
	WithInput(input []byte) WidgetBuilder
	WithViewport(viewport viewports.Viewport) WidgetBuilder
	Now() (Widget, error)
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
