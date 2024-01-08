package dashboards

import (
	"errors"

	"steve.care/network/domain/dashboards"
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/dashboards/widgets/viewports"
)

type builder struct {
	dashboard dashboards.Dashboard
	widget    widgets.Widget
	viewport  viewports.Viewport
}

func createBuilder() Builder {
	out := builder{
		dashboard: nil,
		widget:    nil,
		viewport:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithDashboard adds a dashboard to the builder
func (app *builder) WithDashboard(dashboard dashboards.Dashboard) Builder {
	app.dashboard = dashboard
	return app
}

// WithWidget adds a widget to the builder
func (app *builder) WithWidget(widget widgets.Widget) Builder {
	app.widget = widget
	return app
}

// WithViewport adds a viewport to the builder
func (app *builder) WithViewport(viewport viewports.Viewport) Builder {
	app.viewport = viewport
	return app
}

// Now builds a new Dashboard instance
func (app *builder) Now() (Dashboard, error) {
	if app.dashboard != nil {
		return createDashboardWithDashboard(app.dashboard), nil
	}

	if app.widget != nil {
		return createDashboardWithWidget(app.widget), nil
	}

	if app.viewport != nil {
		return createDashboardWithViewport(app.viewport), nil
	}

	return nil, errors.New("the Dashboard is invalid")
}
