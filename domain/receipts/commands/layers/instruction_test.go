package layers

import (
	"reflect"
	"testing"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers/links"
)

func TestInstruction_isStop_Success(t *testing.T) {
	instruction := NewInstructionWithStopForTests()
	if !instruction.IsStop() {
		t.Errorf("the instruction was expected to contain a stop")
		return
	}

	if instruction.IsRaiseError() {
		t.Errorf("the instruction was expected to NOT contain a raiseError")
		return
	}

	if instruction.IsCondition() {
		t.Errorf("the instruction was expected to NOT contain a condition")
		return
	}

	if instruction.IsAssignment() {
		t.Errorf("the instruction was expected to NOT contain an assignment")
		return
	}

	if instruction.IsLink() {
		t.Errorf("the instruction was expected to NOT contain a link")
		return
	}

	if instruction.IsLayer() {
		t.Errorf("the instruction was expected to NOT contain a layer")
		return
	}
}

func TestInstruction_withRaiseError_Success(t *testing.T) {
	code := uint(56)
	instruction := NewInstructionWithRaiseErrorForTests(code)
	if instruction.IsStop() {
		t.Errorf("the instruction was expected to NOT contain a stop")
		return
	}

	if !instruction.IsRaiseError() {
		t.Errorf("the instruction was expected to contain a raiseError")
		return
	}

	if instruction.IsCondition() {
		t.Errorf("the instruction was expected to NOT contain a condition")
		return
	}

	if instruction.IsAssignment() {
		t.Errorf("the instruction was expected to NOT contain an assignment")
		return
	}

	if instruction.IsLink() {
		t.Errorf("the instruction was expected to NOT contain a link")
		return
	}

	if instruction.IsLayer() {
		t.Errorf("the instruction was expected to NOT contain a layer")
		return
	}

	retRaiseError := instruction.RaiseError()
	if code != retRaiseError {
		t.Errorf("the raisedError code was expected to be %d, %d returned", code, retRaiseError)
		return
	}
}

func TestInstruction_withCondition_Success(t *testing.T) {
	condition := NewConditionForTest(
		"myName",
		NewInstructionsForTests([]Instruction{
			NewInstructionWithStopForTests(),
		}),
	)

	instruction := NewInstructionWithConditionForTests(condition)
	if instruction.IsStop() {
		t.Errorf("the instruction was expected to NOT contain a stop")
		return
	}

	if instruction.IsRaiseError() {
		t.Errorf("the instruction was expected to NOT contain a raiseError")
		return
	}

	if !instruction.IsCondition() {
		t.Errorf("the instruction was expected to contain a condition")
		return
	}

	if instruction.IsAssignment() {
		t.Errorf("the instruction was expected to NOT contain an assignment")
		return
	}

	if instruction.IsLink() {
		t.Errorf("the instruction was expected to NOT contain a link")
		return
	}

	if instruction.IsLayer() {
		t.Errorf("the instruction was expected to NOT contain a layer")
		return
	}

	retCondition := instruction.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the returned condition is invalid")
		return
	}
}

func TestInstruction_withAssignment_Success(t *testing.T) {
	assignment := NewAssignmentForTests(
		"myName",
		NewAssignableWithBytesForTests(NewBytesWithJoinForTests(
			NewBytesReferencesForTests(
				[]BytesReference{
					NewBytesReferenceWithVariableForTests("myVariable"),
					NewBytesReferenceWithBytesForTests([]byte("this is some bytes")),
				},
			),
		)),
	)

	instruction := NewInstructionWithAssignmentForTests(assignment)
	if instruction.IsStop() {
		t.Errorf("the instruction was expected to NOT contain a stop")
		return
	}

	if instruction.IsRaiseError() {
		t.Errorf("the instruction was expected to NOT contain a raiseError")
		return
	}

	if instruction.IsCondition() {
		t.Errorf("the instruction was expected to NOT contain a condition")
		return
	}

	if !instruction.IsAssignment() {
		t.Errorf("the instruction was expected to contain an assignment")
		return
	}

	if instruction.IsLink() {
		t.Errorf("the instruction was expected to NOT contain a link")
		return
	}

	if instruction.IsLayer() {
		t.Errorf("the instruction was expected to NOT contain a layer")
		return
	}

	retAssignment := instruction.Assignment()
	if !reflect.DeepEqual(assignment, retAssignment) {
		t.Errorf("the returned assignment is invalid")
		return
	}
}

func TestInstruction_withLinkInstruction_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	linkInstruction := NewLinkInstructionWithSaveForTests(
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
	)

	instruction := NewInstructionWithLinkInstructionForTests(linkInstruction)
	if instruction.IsStop() {
		t.Errorf("the instruction was expected to NOT contain a stop")
		return
	}

	if instruction.IsRaiseError() {
		t.Errorf("the instruction was expected to NOT contain a raiseError")
		return
	}

	if instruction.IsCondition() {
		t.Errorf("the instruction was expected to NOT contain a condition")
		return
	}

	if instruction.IsAssignment() {
		t.Errorf("the instruction was expected to NOT contain an assignment")
		return
	}

	if !instruction.IsLink() {
		t.Errorf("the instruction was expected to contain a link")
		return
	}

	if instruction.IsLayer() {
		t.Errorf("the instruction was expected to NOT contain a layer")
		return
	}

	retLink := instruction.Link()
	if !reflect.DeepEqual(linkInstruction, retLink) {
		t.Errorf("the returned linkInstruction is invalid")
		return
	}
}

func TestInstruction_withLayerInstruction_Success(t *testing.T) {
	layerInstruction := NewLayerInstructionWithSaveForTests(
		NewLayerForTests(
			NewInstructionsForTests([]Instruction{
				NewInstructionWithStopForTests(),
			}),
			NewOutputForTests(
				"myVariable",
				NewKindWithContinueForTests(),
			),
		),
	)

	instruction := NewInstructionWithLayerInstructionForTests(layerInstruction)
	if instruction.IsStop() {
		t.Errorf("the instruction was expected to NOT contain a stop")
		return
	}

	if instruction.IsRaiseError() {
		t.Errorf("the instruction was expected to NOT contain a raiseError")
		return
	}

	if instruction.IsCondition() {
		t.Errorf("the instruction was expected to NOT contain a condition")
		return
	}

	if instruction.IsAssignment() {
		t.Errorf("the instruction was expected to NOT contain an assignment")
		return
	}

	if instruction.IsLink() {
		t.Errorf("the instruction was expected to NOT contain a link")
		return
	}

	if !instruction.IsLayer() {
		t.Errorf("the instruction was expected to contain a layer")
		return
	}

	retLayer := instruction.Layer()
	if !reflect.DeepEqual(layerInstruction, retLayer) {
		t.Errorf("the returned layerInstruction is invalid")
		return
	}
}

func TestInstruction_withoutParam_returnsError(t *testing.T) {
	_, err := NewInstructionBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
