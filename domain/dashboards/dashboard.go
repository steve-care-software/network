package dashboards

import (
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/hash"
)

type dashboard struct {
	hash    hash.Hash
	title   string
	widgets widgets.Widgets
}

func createDashboard(
	hash hash.Hash,
	title string,
	widgets widgets.Widgets,
) Dashboard {
	out := dashboard{
		hash:    hash,
		title:   title,
		widgets: widgets,
	}

	return &out
}

// Hash returns the hash
func (obj *dashboard) Hash() hash.Hash {
	return obj.hash
}

// Title returns the title
func (obj *dashboard) Title() string {
	return obj.title
}

// Widgets returns the widgets
func (obj *dashboard) Widgets() widgets.Widgets {
	return obj.widgets
}
