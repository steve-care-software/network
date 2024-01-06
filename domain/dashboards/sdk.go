package dashboards

import (
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/hash"
)

// Dashboards represents dashboards
type Dashboards interface {
	Hash() hash.Hash
	List() []Dashboard
}

// Dashboard represents a dashboard
type Dashboard interface {
	Hash() hash.Hash
	Widgets() widgets.Widgets
}
