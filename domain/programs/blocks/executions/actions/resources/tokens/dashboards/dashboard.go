package dashboards

import (
	"steve.care/network/domain/dashboards"
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
)

type dashboard struct {
	dashboard dashboards.Dashboard
	widget    widgets.Widget
	viewport  viewports.Viewport
}

func createDashboardWithDashboard(
	dashboardIns dashboards.Dashboard,
) Dashboard {
	return createDashboardInternally(dashboardIns, nil, nil)
}

func createDashboardWithWidget(
	widget widgets.Widget,
) Dashboard {
	return createDashboardInternally(nil, widget, nil)
}

func createDashboardWithViewport(
	viewport viewports.Viewport,
) Dashboard {
	return createDashboardInternally(nil, nil, viewport)
}

func createDashboardInternally(
	dashboardIns dashboards.Dashboard,
	widget widgets.Widget,
	viewport viewports.Viewport,
) Dashboard {
	out := dashboard{
		dashboard: dashboardIns,
		widget:    widget,
		viewport:  viewport,
	}

	return &out
}

// Hash returns the hash
func (obj *dashboard) Hash() hash.Hash {
	if obj.IsDashboard() {
		return obj.dashboard.Hash()
	}

	if obj.IsWidget() {
		return obj.widget.Hash()
	}

	return obj.viewport.Hash()
}

// IsDashboard returns true if there is a dashboard, false otherwise
func (obj *dashboard) IsDashboard() bool {
	return obj.dashboard != nil
}

// Dashboard returns the dashboard, if any
func (obj *dashboard) Dashboard() dashboards.Dashboard {
	return obj.dashboard
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
