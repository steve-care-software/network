package dashboards

import "steve.care/network/domain/dashboards/widgets"

// NewDashboardsForTests creates a new dashboards for tests
func NewDashboardsForTests(list []Dashboard) Dashboards {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDashboardForTests creates a new dashboard for tests
func NewDashboardForTests(title string, widgets widgets.Widgets) Dashboard {
	ins, err := NewDashboardBuilder().Create().WithTitle(title).WithWidgets(widgets).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
