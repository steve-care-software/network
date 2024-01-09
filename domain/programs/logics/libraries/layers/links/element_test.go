package links

import (
	"bytes"
	"reflect"
	"testing"

	"steve.care/network/domain/hash"
)

func TestElement_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	element := NewElementForTests(*pLayer)
	retLayer := element.Layer()
	if !bytes.Equal(pLayer.Bytes(), retLayer.Bytes()) {
		t.Errorf("the returned layer is invalid")
		return
	}

	if element.HasCondition() {
		t.Errorf("the element was expected to NOT contain condition")
		return
	}
}

func TestElement_withCondition_Success(t *testing.T) {
	condition := NewConditionForTests(
		NewConditionResourceForTests(23),
	)

	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	element := NewElementWithConditionForTests(*pLayer, condition)
	retLayer := element.Layer()
	if !bytes.Equal(pLayer.Bytes(), retLayer.Bytes()) {
		t.Errorf("the returned layer is invalid")
		return
	}

	if !element.HasCondition() {
		t.Errorf("the element was expected to contain condition")
		return
	}

	retCondition := element.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the condition is invalid")
		return
	}
}
