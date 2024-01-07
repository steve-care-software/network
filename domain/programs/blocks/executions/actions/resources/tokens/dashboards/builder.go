package dashboards

import (
	"errors"

	"steve.care/network/domain/dashboards"
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/dashboards/widgets/viewports/dimensions"
	"steve.care/network/domain/dashboards/widgets/viewports/positions"
)

type builder struct {
	dashboard dashboards.Dashboard
	widgets   widgets.Widgets
	widget    widgets.Widget
	viewport  viewports.Viewport
	dimension dimensions.Dimension
	position  positions.Position
}

func createBuilder() Builder {
	out := builder{
		dashboard: nil,
		widgets:   nil,
		widget:    nil,
		viewport:  nil,
		dimension: nil,
		position:  nil,
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

// WithWidgets adds a widgets to the builder
func (app *builder) WithWidgets(widgets widgets.Widgets) Builder {
	app.widgets = widgets
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

// WithDimension adds a dimension to the builder
func (app *builder) WithDimension(dimension dimensions.Dimension) Builder {
	app.dimension = dimension
	return app
}

// WithPosition adds a position to the builder
func (app *builder) WithPosition(position positions.Position) Builder {
	app.position = position
	return app
}

// Now builds a new Dashboard instance
func (app *builder) Now() (Dashboard, error) {
	if app.dashboard != nil {
		return createDashboardWithDashboard(app.dashboard), nil
	}

	if app.widgets != nil {
		return createDashboardWithWidgets(app.widgets), nil
	}

	if app.widget != nil {
		return createDashboardWithWidget(app.widget), nil
	}

	if app.viewport != nil {
		return createDashboardWithViewport(app.viewport), nil
	}

	if app.dimension != nil {
		return createDashboardWithDimension(app.dimension), nil
	}

	if app.position != nil {
		return createDashboardWithPosition(app.position), nil
	}

	return nil, errors.New("the Dashboard is invalid")
}