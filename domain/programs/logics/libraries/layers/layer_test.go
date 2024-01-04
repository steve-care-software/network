package layers

import (
	"reflect"
	"testing"
)

func TestLayer_Success(t *testing.T) {
	instructions := NewInstructionsForTests([]Instruction{
		NewInstructionWithStopForTests(),
	})

	output := NewOutputForTests(
		"myVariable",
		NewKindWithContinueForTests(),
	)

	layer := NewLayerForTests(
		instructions,
		output,
	)

	retInstructions := layer.Instructions()
	if !reflect.DeepEqual(instructions, retInstructions) {
		t.Errorf("the returned instructions is invalid")
		return
	}

	retOutput := layer.Output()
	if !reflect.DeepEqual(output, retOutput) {
		t.Errorf("the returned output is invalid")
		return
	}

	if layer.HasInput() {
		t.Errorf("the layer was expected to NOT contain input")
		return
	}
}

func TestLayer_withInput_Success(t *testing.T) {
	instructions := NewInstructionsForTests([]Instruction{
		NewInstructionWithStopForTests(),
	})

	output := NewOutputForTests(
		"myVariable",
		NewKindWithContinueForTests(),
	)

	input := "myInput"
	layer := NewLayerWithInputForTests(
		instructions,
		output,
		input,
	)

	retInstructions := layer.Instructions()
	if !reflect.DeepEqual(instructions, retInstructions) {
		t.Errorf("the returned instructions is invalid")
		return
	}

	retOutput := layer.Output()
	if !reflect.DeepEqual(output, retOutput) {
		t.Errorf("the returned output is invalid")
		return
	}

	if !layer.HasInput() {
		t.Errorf("the layer was expected to contain input")
		return
	}

	retInput := layer.Input()
	if !reflect.DeepEqual(input, retInput) {
		t.Errorf("the returned input is invalid")
		return
	}
}
