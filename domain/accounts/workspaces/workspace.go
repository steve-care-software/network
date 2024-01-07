package workspaces

import "steve.care/network/domain/dashboards"

type workspace struct {
	dashboards dashboards.Dashboards
	root       dashboards.Dashboard
}

func createWorkspace(
	dashboards dashboards.Dashboards,
	root dashboards.Dashboard,
) Workspace {
	out := workspace{
		dashboards: dashboards,
		root:       root,
	}

	return &out
}

// Dashboards returns the dashboards
func (obj *workspace) Dashboards() dashboards.Dashboards {
	return obj.dashboards
}

// Root returns the root dashboard
func (obj *workspace) Root() dashboards.Dashboard {
	return obj.root
}
