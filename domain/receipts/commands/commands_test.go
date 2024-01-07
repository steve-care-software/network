package commands

import (
	"reflect"
	"testing"

	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/receipts/commands/results"
)

func TestCommands_Success(t *testing.T) {
	list := []Command{
		NewCommandForTests(
			[]byte("this is the command input"),
			layers.NewLayerForTests(
				layers.NewInstructionsForTests([]layers.Instruction{
					layers.NewInstructionWithStopForTests(),
				}),
				layers.NewOutputForTests(
					"myVariable",
					layers.NewKindWithContinueForTests(),
				),
			),
			results.NewResultWithSuccessForTests(
				results.NewSuccessForTests(
					[]byte("this is some bytes"),
					layers.NewKindWithPromptForTests(),
				),
			),
		),
	}

	ins := NewCommandsForTests(list)

	retList := ins.List()
	if !reflect.DeepEqual(list, retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestCommands_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCommands_withEmptyList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithList([]Command{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
