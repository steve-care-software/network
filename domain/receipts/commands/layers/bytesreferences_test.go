package layers

import "testing"

func TestBytesReferences_withList_Success(t *testing.T) {
	list := []BytesReference{
		NewBytesReferenceWithVariableForTests("myVariable"),
		NewBytesReferenceWithBytesForTests([]byte("this is some bytes")),
	}

	ins := NewBytesReferencesForTests(list)
	retList := ins.List()
	if len(list) != len(retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestBytesReferences_withEmptyList_returnsError(t *testing.T) {
	list := []BytesReference{}
	_, err := NewBytesReferencesBuilder().Create().WithList(list).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}

func TestBytesReferences_withoutList_returnsError(t *testing.T) {
	_, err := NewBytesReferencesBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}
