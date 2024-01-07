package widgets

import (
	"bytes"
	"reflect"
	"testing"

	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
)

func TestWidget_Success(t *testing.T) {
	title := "this is a title"
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	input := []byte("this is an input")
	ins := NewWidgetForTests(title, *pProgram, input)

	retTitle := ins.Title()
	if title != retTitle {
		t.Errorf("the title was expected to be '%s', '%s' returned", title, retTitle)
		return
	}

	retProgram := ins.Program()
	if !bytes.Equal(pProgram.Bytes(), retProgram.Bytes()) {
		t.Errorf("the program was expected to be '%s', '%s' returned", pProgram.String(), retProgram.String())
		return
	}

	retInput := ins.Input()
	if !bytes.Equal(input, retInput) {
		t.Errorf("the returned input is invalid")
		return
	}

	if ins.HasViewport() {
		t.Errorf("the widget was expected to NOT contain a Viewport")
		return
	}
}

func TestWidget_withViewport_Success(t *testing.T) {
	title := "this is a title"
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	input := []byte("this is an input")
	viewport := viewports.NewViewportForTests(uint(2), uint(50))
	ins := NewWidgetWithViewportForTests(title, *pProgram, input, viewport)

	retTitle := ins.Title()
	if title != retTitle {
		t.Errorf("the title was expected to be '%s', '%s' returned", title, retTitle)
		return
	}

	retProgram := ins.Program()
	if !bytes.Equal(pProgram.Bytes(), retProgram.Bytes()) {
		t.Errorf("the program was expected to be '%s', '%s' returned", pProgram.String(), retProgram.String())
		return
	}

	retInput := ins.Input()
	if !bytes.Equal(input, retInput) {
		t.Errorf("the returned input is invalid")
		return
	}

	if !ins.HasViewport() {
		t.Errorf("the widget was expected to contain a Viewport")
		return
	}

	retViewport := ins.Viewport()
	if !reflect.DeepEqual(viewport, retViewport) {
		t.Errorf("the returned viewport is invalid")
		return
	}
}

func TestWidget_withoutTitle_returnsError(t *testing.T) {
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	input := []byte("this is an input")
	_, err := NewWidgetBuilder().Create().WithProgram(*pProgram).WithInput(input).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestWidget_withoutProgram_returnsError(t *testing.T) {
	title := "this is a title"
	input := []byte("this is an input")
	_, err := NewWidgetBuilder().Create().WithTitle(title).WithInput(input).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestWidget_withEmptyInput_returnsError(t *testing.T) {
	title := "this is a title"
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	input := []byte{}
	_, err := NewWidgetBuilder().Create().WithTitle(title).WithProgram(*pProgram).WithInput(input).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestWidget_withoutInput_returnsError(t *testing.T) {
	title := "this is a title"
	pProgram, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	_, err := NewWidgetBuilder().Create().WithTitle(title).WithProgram(*pProgram).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
