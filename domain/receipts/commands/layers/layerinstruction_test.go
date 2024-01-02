package layers

import (
	"bytes"
	"reflect"
	"testing"

	"steve.care/network/domain/hash"
)

func TestLayerInstruction_withSave_Success(t *testing.T) {
	layer := NewLayerForTests(
		NewInstructionsForTests([]Instruction{
			NewInstructionWithStopForTests(),
		}),
		NewOutputForTests(
			"myVariable",
			NewKindWithContinueForTests(),
		),
	)

	layerInstruction := NewLayerInstructionWithSaveForTests(layer)
	if !layerInstruction.IsSave() {
		t.Errorf("the layerInstruction was expected to contain a save")
		return
	}

	if layerInstruction.IsDelete() {
		t.Errorf("the layerInstruction was expected to NOT contain a delete")
		return
	}

	retSave := layerInstruction.Save()
	if !reflect.DeepEqual(layer, retSave) {
		t.Errorf("the returned save layer is invalid")
		return
	}
}

func TestLayerInstruction_withDelete_Success(t *testing.T) {
	pDel, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	layerInstruction := NewLayerInstructionWithDeleteForTests(*pDel)
	if layerInstruction.IsSave() {
		t.Errorf("the layerInstruction was expected to NOT contain a save")
		return
	}

	if !layerInstruction.IsDelete() {
		t.Errorf("the layerInstruction was expected to contain a delete")
		return
	}

	retDel := layerInstruction.Delete()
	if !bytes.Equal(pDel.Bytes(), retDel.Bytes()) {
		t.Errorf("the returned delete layer hash is invalid")
		return
	}
}

func TestLayerInstruction_withoutParam_returnsError(t *testing.T) {
	_, err := NewLayerInstructionBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
