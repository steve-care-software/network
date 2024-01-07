package dashboards

import (
	"steve.care/network/domain/dashboards"
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/dashboards/widgets/viewports/dimensions"
	"steve.care/network/domain/dashboards/widgets/viewports/positions"
	"steve.care/network/domain/hash"
)

type dashboard struct {
	dashboard dashboards.Dashboard
	widgets   widgets.Widgets
	widget    widgets.Widget
	viewport  viewports.Viewport
	dimension dimensions.Dimension
	position  positions.Position
}

func createDashboardWithDashboard(
	dashboardIns dashboards.Dashboard,
) Dashboard {
	return createDashboardInternally(dashboardIns, nil, nil, nil, nil, nil)
}

func createDashboardWithWidgets(
	widgets widgets.Widgets,
) Dashboard {
	return createDashboardInternally(nil, widgets, nil, nil, nil, nil)
}

func createDashboardWithWidget(
	widget widgets.Widget,
) Dashboard {
	return createDashboardInternally(nil, nil, widget, nil, nil, nil)
}

func createDashboardWithViewport(
	viewport viewports.Viewport,
) Dashboard {
	return createDashboardInternally(nil, nil, nil, viewport, nil, nil)
}

func createDashboardWithDimension(
	dimension dimensions.Dimension,
) Dashboard {
	return createDashboardInternally(nil, nil, nil, nil, dimension, nil)
}

func createDashboardWithPosition(
	position positions.Position,
) Dashboard {
	return createDashboardInternally(nil, nil, nil, nil, nil, position)
}

func createDashboardInternally(
	dashboardIns dashboards.Dashboard,
	widgets widgets.Widgets,
	widget widgets.Widget,
	viewport viewports.Viewport,
	dimension dimensions.Dimension,
	position positions.Position,
) Dashboard {
	out := dashboard{
		dashboard: dashboardIns,
		widgets:   widgets,
		widget:    widget,
		viewport:  viewport,
		dimension: dimension,
		position:  position,
	}

	return &out
}

// Hash returns the hash
func (obj *dashboard) Hash() hash.Hash {
	if obj.IsDashboard() {
		return obj.dashboard.Hash()
	}

	if obj.IsWidgets() {
		return obj.widgets.Hash()
	}

	if obj.IsWidget() {
		return obj.widget.Hash()
	}

	if obj.IsViewport() {
		return obj.viewport.Hash()
	}

	if obj.IsDimension() {
		return obj.dimension.Hash()
	}

	return obj.position.Hash()
}

// IsDashboard returns true if there is a dashboard, false otherwise
func (obj *dashboard) IsDashboard() bool {
	return obj.dashboard != nil
}

// Dashboard returns the dashboard, if any
func (obj *dashboard) Dashboard() dashboards.Dashboard {
	return obj.dashboard
}

// IsWidgets returns true if there is a widgets, false otherwise
func (obj *dashboard) IsWidgets() bool {
	return obj.widgets != nil
}

// Widgets returns the dashboard, if any
func (obj *dashboard) Widgets() widgets.Widgets {
	return obj.widgets
}

// IsWidget returns true if there is a widget, false otherwise
func (obj *dashboard) IsWidget() bool {
	return obj.widget != nil
}

// Widget returns the widget, if any
func (obj *dashboard) Widget() widgets.Widget {
	return obj.widget
}

// IsViewport returns true if there is a viewport, false otherwise
func (obj *dashboard) IsViewport() bool {
	return obj.viewport != nil
}

// Viewport returns the viewport, if any
func (obj *dashboard) Viewport() viewports.Viewport {
	return obj.viewport
}

// IsDimension returns true if there is a dimension, false otherwise
func (obj *dashboard) IsDimension() bool {
	return obj.dimension != nil
}

// Dimension returns the dimension, if any
func (obj *dashboard) Dimension() dimensions.Dimension {
	return obj.dimension
}

// IsPosition returns true if there is a position, false otherwise
func (obj *dashboard) IsPosition() bool {
	return obj.position != nil
}

// Position returns the position, if any
func (obj *dashboard) Position() positions.Position {
	return obj.position
}
