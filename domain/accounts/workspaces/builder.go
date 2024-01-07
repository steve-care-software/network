package workspaces

import (
	"errors"

	"steve.care/network/domain/dashboards"
)

type builder struct {
	dashboards dashboards.Dashboards
	root       dashboards.Dashboard
}

func createBuilder() Builder {
	out := builder{
		dashboards: nil,
		root:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithDashboards add dashboards to the builder
func (app *builder) WithDashboards(dashboards dashboards.Dashboards) Builder {
	app.dashboards = dashboards
	return app
}

// WithRoot add root dashboard to the builder
func (app *builder) WithRoot(root dashboards.Dashboard) Builder {
	app.root = root
	return app
}

// Now builds a new Workspace instance
func (app *builder) Now() (Workspace, error) {
	if app.dashboards == nil {
		return nil, errors.New("the dashboards is mandatory in order to build a Workspace instance")
	}

	if app.root == nil {
		return nil, errors.New("the root dashboard is mandatory in order to build a Workspace instance")
	}

	return createWorkspace(app.dashboards, app.root), nil
}
