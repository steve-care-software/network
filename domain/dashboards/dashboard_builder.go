package dashboards

import (
	"errors"

	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/hash"
)

type dashboardBuilder struct {
	hashAdapter hash.Adapter
	title       string
	widgets     widgets.Widgets
}

func createDashboardBuilder(
	hashAdapter hash.Adapter,
) DashboardBuilder {
	out := dashboardBuilder{
		hashAdapter: hashAdapter,
		title:       "",
		widgets:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *dashboardBuilder) Create() DashboardBuilder {
	return createDashboardBuilder(
		app.hashAdapter,
	)
}

// WithTitle adds a title to the builder
func (app *dashboardBuilder) WithTitle(title string) DashboardBuilder {
	app.title = title
	return app
}

// WithWidgets add widgets to the builder
func (app *dashboardBuilder) WithWidgets(widgets widgets.Widgets) DashboardBuilder {
	app.widgets = widgets
	return app
}

// Now builds a new Dashboard instance
func (app *dashboardBuilder) Now() (Dashboard, error) {
	if app.title == "" {
		return nil, errors.New("the title is mandatory in order to build a Dashboard instance")
	}

	if app.widgets == nil {
		return nil, errors.New("the widgets is mandatory in order to build a Dashboard instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.title),
		app.widgets.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createDashboard(*pHash, app.title, app.widgets), nil
}
