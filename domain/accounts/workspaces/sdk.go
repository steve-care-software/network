package workspaces

import (
	"steve.care/network/domain/dashboards"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the workspace builder
type Builder interface {
	Create() Builder
	WithDashboards(dashboards dashboards.Dashboards) Builder
	WithRoot(root dashboards.Dashboard) Builder
	Now() (Workspace, error)
}

// Workspace represents a workspace
type Workspace interface {
	Dashboards() dashboards.Dashboards
	Root() dashboards.Dashboard
}
