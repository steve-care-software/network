package dashboards

import (
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/hash"
)

// Builder represents the dashboards builder
type Builder interface {
	Create() Builder
	WithList(list []Dashboard) Builder
	Now() (Dashboards, error)
}

// Dashboards represents dashboards
type Dashboards interface {
	Hash() hash.Hash
	List() []Dashboard
}

// DashboardBuilder represents the dashboard builder
type DashboardBuilder interface {
	Create() DashboardBuilder
	WithTitle(title string) DashboardBuilder
	WithWidgets(widgets widgets.Widgets) DashboardBuilder
	Now() (Dashboard, error)
}

// Dashboard represents a dashboard
type Dashboard interface {
	Hash() hash.Hash
	Title() string
	Widgets() widgets.Widgets
}
