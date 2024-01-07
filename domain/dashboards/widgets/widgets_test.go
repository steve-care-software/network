package widgets

import (
	"reflect"
	"testing"

	"steve.care/network/domain/hash"
)

func TestWidgets_Success(t *testing.T) {
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	list := []Widget{
		NewWidgetForTests(
			"this is a title",
			*pProgram,
			[]byte("this is an input"),
		),
	}

	ins := NewWidgetsForTests(list)

	retList := ins.List()
	if !reflect.DeepEqual(list, retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestWidgets_withEmptyList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithList([]Widget{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestWidgets_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
