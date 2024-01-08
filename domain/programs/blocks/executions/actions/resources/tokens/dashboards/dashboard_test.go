package dashboards

import (
	"reflect"
	"testing"

	"steve.care/network/domain/dashboards"
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
)

func TestDashboard_withDashboard_Success(t *testing.T) {
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	dashoard := dashboards.NewDashboardForTests(
		"this is a title",
		widgets.NewWidgetsForTests([]widgets.Widget{
			widgets.NewWidgetForTests(
				"this is a title",
				*pProgram,
				[]byte("this is an input"),
			),
		}),
	)

	ins := NewDashboardWithDashboardForTests(dashoard)

	if !ins.IsDashboard() {
		t.Errorf("the dashboard was expected to contain a dashboard")
		return
	}

	if ins.IsWidgets() {
		t.Errorf("the dashboard was expected to NOT contain a widgets")
		return
	}

	if ins.IsWidget() {
		t.Errorf("the dashboard was expected to NOT contain a widget")
		return
	}

	if ins.IsViewport() {
		t.Errorf("the dashboard was expected to NOT contain a viewport")
		return
	}

	retDashboard := ins.Dashboard()
	if !reflect.DeepEqual(dashoard, retDashboard) {
		t.Errorf("the returned dashboard is invalid")
		return
	}
}

func TestDashboard_withWidgets_Success(t *testing.T) {
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	widgets := widgets.NewWidgetsForTests([]widgets.Widget{
		widgets.NewWidgetForTests(
			"this is a title",
			*pProgram,
			[]byte("this is an input"),
		),
	})

	ins := NewDashboardWithWidgetsForTests(widgets)

	if ins.IsDashboard() {
		t.Errorf("the dashboard was expected to NOT contain a dashboard")
		return
	}

	if !ins.IsWidgets() {
		t.Errorf("the dashboard was expected to contain a widgets")
		return
	}

	if ins.IsWidget() {
		t.Errorf("the dashboard was expected to NOT contain a widget")
		return
	}

	if ins.IsViewport() {
		t.Errorf("the dashboard was expected to NOT contain a viewport")
		return
	}

	retWidgets := ins.Widgets()
	if !reflect.DeepEqual(widgets, retWidgets) {
		t.Errorf("the returned widgets is invalid")
		return
	}
}

func TestDashboard_withWidget_Success(t *testing.T) {
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	widget := widgets.NewWidgetForTests(
		"this is a title",
		*pProgram,
		[]byte("this is an input"),
	)

	ins := NewDashboardWithWidgetForTests(widget)

	if ins.IsDashboard() {
		t.Errorf("the dashboard was expected to NOT contain a dashboard")
		return
	}

	if ins.IsWidgets() {
		t.Errorf("the dashboard was expected to NOT contain a widgets")
		return
	}

	if !ins.IsWidget() {
		t.Errorf("the dashboard was expected to contain a widget")
		return
	}

	if ins.IsViewport() {
		t.Errorf("the dashboard was expected to NOT contain a viewport")
		return
	}

	retWidget := ins.Widget()
	if !reflect.DeepEqual(widget, retWidget) {
		t.Errorf("the returned widget is invalid")
		return
	}
}

func TestDashboard_withViewport_Success(t *testing.T) {
	viewport := viewports.NewViewportForTests(uint(0), uint(45))
	ins := NewDashboardWithViewportForTests(viewport)

	if ins.IsDashboard() {
		t.Errorf("the dashboard was expected to NOT contain a dashboard")
		return
	}

	if ins.IsWidgets() {
		t.Errorf("the dashboard was expected to NOT contain a widgets")
		return
	}

	if ins.IsWidget() {
		t.Errorf("the dashboard was expected to NOT contain a widget")
		return
	}

	if !ins.IsViewport() {
		t.Errorf("the dashboard was expected to contain a viewport")
		return
	}

	retViewport := ins.Viewport()
	if !reflect.DeepEqual(viewport, retViewport) {
		t.Errorf("the returned viewport is invalid")
		return
	}
}

func TestDashboard_withoutParam_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
