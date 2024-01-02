package layers

import (
	"bytes"
	"reflect"
	"testing"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers/links"
)

func TestLinkInstruction_withSave_Success(t *testing.T) {
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

	linkInstruction := NewLinkInstructionWithSaveForTests(link)
	if !linkInstruction.IsSave() {
		t.Errorf("the linkInstruction was expected to contain a save")
		return
	}

	if linkInstruction.IsDelete() {
		t.Errorf("the linkInstruction was expected to NOT contain a delete")
		return
	}

	retSave := linkInstruction.Save()
	if !reflect.DeepEqual(link, retSave) {
		t.Errorf("the returned save link is invalid")
		return
	}
}

func TestLinkInstruction_withDelete_Success(t *testing.T) {
	pDel, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	linkInstruction := NewLinkInstructionWithDeleteForTests(*pDel)
	if linkInstruction.IsSave() {
		t.Errorf("the linkInstruction was expected to NOT contain a save")
		return
	}

	if !linkInstruction.IsDelete() {
		t.Errorf("the linkInstruction was expected to contain a delete")
		return
	}

	retDel := linkInstruction.Delete()
	if !bytes.Equal(pDel.Bytes(), retDel.Bytes()) {
		t.Errorf("the returned delete link hash is invalid")
		return
	}
}

func TestLinkInstruction_withoutParam_returnsError(t *testing.T) {
	_, err := NewLinkInstructionBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
