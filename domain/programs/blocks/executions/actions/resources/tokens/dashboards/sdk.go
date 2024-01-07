package dashboards

import (
	"steve.care/network/domain/dashboards"
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a dashboard builder
type Builder interface {
	Create() Builder
	WithDashboard(dashboard dashboards.Dashboard) Builder
	WithWidgets(widgets widgets.Widgets) Builder
	WithWidget(widget widgets.Widget) Builder
	WithViewport(viewport viewports.Viewport) Builder
	Now() (Dashboard, error)
}

// Dashboard represents a dashboard resource
type Dashboard interface {
	Hash() hash.Hash
	IsDashboard() bool
	Dashboard() dashboards.Dashboard
	IsWidgets() bool
	Widgets() widgets.Widgets
	IsWidget() bool
	Widget() widgets.Widget
	IsViewport() bool
	Viewport() viewports.Viewport
}
