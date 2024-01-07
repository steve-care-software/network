package commands

import (
	"reflect"
	"testing"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
	"steve.care/network/domain/receipts/commands/results"
)

func TestLink_Success(t *testing.T) {
	input := []byte("this is an input")
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))

	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	link := links.NewLinkForTests(
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
	)

	command := NewCommandForTests(
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
	)

	ins := NewLinkForTests(input, link, command)

	retInput := ins.Input()
	if !reflect.DeepEqual(input, retInput) {
		t.Errorf("the returned input is invalid")
		return
	}

	retLink := ins.Link()
	if !reflect.DeepEqual(link, retLink) {
		t.Errorf("the returned link is invalid")
		return
	}

	retCommand := ins.Command()
	if !reflect.DeepEqual(command, retCommand) {
		t.Errorf("the returned command is invalid")
		return
	}
}

func TestLink_withoutInput_returnsError(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))

	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	link := links.NewLinkForTests(
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
	)

	command := NewCommandForTests(
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
	)

	_, err := NewLinkBuilder().Create().WithLink(link).WithCommand(command).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestLink_withoutLink_returnsError(t *testing.T) {
	input := []byte("this is an input")
	command := NewCommandForTests(
		input,
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
	)

	_, err := NewLinkBuilder().Create().WithInput(input).WithCommand(command).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestLink_withoutCommand_returnsError(t *testing.T) {
	input := []byte("this is an input")
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))

	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	link := links.NewLinkForTests(
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
	)

	_, err := NewLinkBuilder().Create().WithLink(link).WithInput(input).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
