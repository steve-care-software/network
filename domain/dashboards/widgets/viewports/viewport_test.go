package viewports

import (
	"testing"
)

func TestViewport_Success(t *testing.T) {
	row := uint(0)
	height := uint(45)
	ins := NewViewportForTests(row, height)

	retRow := ins.Row()
	if row != retRow {
		t.Errorf("the row was expected to be %d, %d returned", row, retRow)
		return
	}
}

func TestViewport_withZeroHeight_returnsError(t *testing.T) {
	row := uint(0)
	height := uint(0)
	_, err := NewBuilder().Create().WithRow(row).WithHeight(height).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestViewport_withoutRow_returnsError(t *testing.T) {
	height := uint(45)
	_, err := NewBuilder().Create().WithHeight(height).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestViewport_withoutHeight_returnsError(t *testing.T) {
	row := uint(0)
	_, err := NewBuilder().Create().WithRow(row).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
