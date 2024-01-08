package dashboards

import (
	"steve.care/network/domain/dashboards"
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/dashboards/widgets/viewports"
)

// NewDashboardWithDashboardForTests creates a new dashboard with dashboard for tests
func NewDashboardWithDashboardForTests(dashboard dashboards.Dashboard) Dashboard {
	ins, err := NewBuilder().Create().WithDashboard(dashboard).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDashboardWithWidgetForTests creates a new dashboard with widget for tests
func NewDashboardWithWidgetForTests(widget widgets.Widget) Dashboard {
	ins, err := NewBuilder().Create().WithWidget(widget).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDashboardWithViewportForTests creates a new dashboard with viewport for tests
func NewDashboardWithViewportForTests(viewport viewports.Viewport) Dashboard {
	ins, err := NewBuilder().Create().WithViewport(viewport).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
