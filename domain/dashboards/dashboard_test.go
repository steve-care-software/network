package dashboards

import (
	"reflect"
	"testing"

	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/hash"
)

func TestDashboard_Success(t *testing.T) {
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))

	widgets := widgets.NewWidgetsForTests([]widgets.Widget{
		widgets.NewWidgetForTests(
			"this is a title",
			*pProgram,
			[]byte("this is an input"),
		),
	})

	title := "this is a title"
	ins := NewDashboardForTests(title, widgets)

	retTitle := ins.Title()
	if title != retTitle {
		t.Errorf("the title was expected to be '%s', '%s' returned", title, retTitle)
		return
	}

	retWidgets := ins.Widgets()
	if !reflect.DeepEqual(widgets, retWidgets) {
		t.Errorf("the returned widgets is invalid")
		return
	}
}

func TestDashboard_withoutTitle_returnsError(t *testing.T) {
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))

	widgets := widgets.NewWidgetsForTests([]widgets.Widget{
		widgets.NewWidgetForTests(
			"this is a title",
			*pProgram,
			[]byte("this is an input"),
		),
	})

	_, err := NewDashboardBuilder().Create().WithWidgets(widgets).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestDashboard_withoutWidgets_returnsError(t *testing.T) {
	title := "this is a title"
	_, err := NewDashboardBuilder().Create().WithTitle(title).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
