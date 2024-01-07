package dashboards

import (
	"reflect"
	"testing"

	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/hash"
)

func TestDashboards_Success(t *testing.T) {
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	list := []Dashboard{
		NewDashboardForTests(
			"this is a title",
			widgets.NewWidgetsForTests([]widgets.Widget{
				widgets.NewWidgetForTests(
					"this is a title",
					*pProgram,
					[]byte("this is an input"),
				),
			}),
		),
	}

	ins := NewDashboardsForTests(list)

	retList := ins.List()
	if !reflect.DeepEqual(list, retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestDashboards_withEmptyList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithList([]Dashboard{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestDashboards_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
