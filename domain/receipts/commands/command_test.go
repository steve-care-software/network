package commands

import (
	"reflect"
	"testing"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
	"steve.care/network/domain/receipts/commands/results"
)

func TestCommand_Success(t *testing.T) {
	input := []byte("this is the command input")
	layer := layers.NewLayerForTests(
		layers.NewInstructionsForTests([]layers.Instruction{
			layers.NewInstructionWithStopForTests(),
		}),
		layers.NewOutputForTests(
			"myVariable",
			layers.NewKindWithContinueForTests(),
		),
	)

	result := results.NewResultWithSuccessForTests(
		results.NewSuccessForTests(
			[]byte("this is some bytes"),
			layers.NewKindWithPromptForTests(),
		),
	)

	ins := NewCommandForTests(
		input,
		layer,
		result,
	)

	retInput := ins.Input()
	if !reflect.DeepEqual(input, retInput) {
		t.Errorf("the returned input is invalid")
		return
	}

	retLayer := ins.Layer()
	if !reflect.DeepEqual(layer, retLayer) {
		t.Errorf("the returned layer is invalid")
		return
	}

	retResult := ins.Result()
	if !reflect.DeepEqual(result, retResult) {
		t.Errorf("the returned result is invalid")
		return
	}

	if ins.HasParent() {
		t.Errorf("the command was expected to NOT contain a parent")
		return
	}
}

func TestCommand_withParent_Success(t *testing.T) {
	input := []byte("this is the command input")
	layer := layers.NewLayerForTests(
		layers.NewInstructionsForTests([]layers.Instruction{
			layers.NewInstructionWithStopForTests(),
		}),
		layers.NewOutputForTests(
			"myVariable",
			layers.NewKindWithContinueForTests(),
		),
	)

	result := results.NewResultWithSuccessForTests(
		results.NewSuccessForTests(
			[]byte("this is some bytes"),
			layers.NewKindWithPromptForTests(),
		),
	)

	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))

	parent := NewLinkForTests(
		[]byte("this is an input"),
		links.NewLinkForTests(
			links.NewOriginForTests(
				links.NewOriginResourceForTests(*pFirstLayer),
				links.NewOperatorWithAndForTests(),
				links.NewOriginValueWithResourceForTests(
					links.NewOriginResourceForTests(*pSecondLayer),
				),
			),
			links.NewElementsForTests([]links.Element{
				links.NewElementForTests(*pLayer),
			}),
		),
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
	)

	ins := NewCommandWithParentForTests(
		input,
		layer,
		result,
		parent,
	)

	retInput := ins.Input()
	if !reflect.DeepEqual(input, retInput) {
		t.Errorf("the returned input is invalid")
		return
	}

	retLayer := ins.Layer()
	if !reflect.DeepEqual(layer, retLayer) {
		t.Errorf("the returned layer is invalid")
		return
	}

	retResult := ins.Result()
	if !reflect.DeepEqual(result, retResult) {
		t.Errorf("the returned result is invalid")
		return
	}

	if !ins.HasParent() {
		t.Errorf("the command was expected to contain a parent")
		return
	}

	retParent := ins.Parent()
	if !reflect.DeepEqual(parent, retParent) {
		t.Errorf("the returned parent Link is invalid")
		return
	}
}

func TestCommand_withoutInput_returnsError(t *testing.T) {
	layer := layers.NewLayerForTests(
		layers.NewInstructionsForTests([]layers.Instruction{
			layers.NewInstructionWithStopForTests(),
		}),
		layers.NewOutputForTests(
			"myVariable",
			layers.NewKindWithContinueForTests(),
		),
	)

	result := results.NewResultWithSuccessForTests(
		results.NewSuccessForTests(
			[]byte("this is some bytes"),
			layers.NewKindWithPromptForTests(),
		),
	)

	_, err := NewCommandBuilder().Create().WithLayer(layer).WithResult(result).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCommand_withEmptyInput_returnsError(t *testing.T) {
	layer := layers.NewLayerForTests(
		layers.NewInstructionsForTests([]layers.Instruction{
			layers.NewInstructionWithStopForTests(),
		}),
		layers.NewOutputForTests(
			"myVariable",
			layers.NewKindWithContinueForTests(),
		),
	)

	result := results.NewResultWithSuccessForTests(
		results.NewSuccessForTests(
			[]byte("this is some bytes"),
			layers.NewKindWithPromptForTests(),
		),
	)

	_, err := NewCommandBuilder().Create().WithInput([]byte{}).WithLayer(layer).WithResult(result).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCommand_withoutLayer_returnsError(t *testing.T) {
	input := []byte("this is the command input")
	result := results.NewResultWithSuccessForTests(
		results.NewSuccessForTests(
			[]byte("this is some bytes"),
			layers.NewKindWithPromptForTests(),
		),
	)

	_, err := NewCommandBuilder().Create().WithInput(input).WithResult(result).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCommand_withoutResult_returnsError(t *testing.T) {
	input := []byte("this is the command input")
	layer := layers.NewLayerForTests(
		layers.NewInstructionsForTests([]layers.Instruction{
			layers.NewInstructionWithStopForTests(),
		}),
		layers.NewOutputForTests(
			"myVariable",
			layers.NewKindWithContinueForTests(),
		),
	)

	_, err := NewCommandBuilder().Create().WithInput(input).WithLayer(layer).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
