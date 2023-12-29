package layers

import (
	"bytes"
	"testing"
)

func TestBytesReference_withVariable_Success(t *testing.T) {
	variable := "myVariable"
	ins := NewBytesReferenceWithVariableForTests(variable)
	if !ins.IsVariable() {
		t.Errorf("the variable was expected to be valid")
		return
	}

	if ins.Variable() != variable {
		t.Errorf("the variable was expected to be '%s', '%s' returned", variable, ins.Variable())
		return
	}

	if ins.IsBytes() {
		t.Errorf("the bytes was expected to NOT be valid")
		return
	}
}

func TestBytesReference_withBytes_Success(t *testing.T) {
	bytesIns := []byte("this is some bytes")
	ins := NewBytesReferenceWithBytesForTests(bytesIns)
	if ins.IsVariable() {
		t.Errorf("the variable was expected to NOT be valid")
		return
	}

	if !ins.IsBytes() {
		t.Errorf("the bytes was expected to be valid")
		return
	}

	if !bytes.Equal(bytesIns, ins.Bytes()) {
		t.Errorf("the returned bytes were expected to be equal")
		return
	}
}

func TestBytesReference_withEmptyBytes_returnsError(t *testing.T) {
	_, err := NewBytesReferenceBuilder().Create().WithBytes([]byte{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesReference_withEmptyVariable_returnsError(t *testing.T) {
	_, err := NewBytesReferenceBuilder().Create().WithVariable("").Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestBytesReference_withoutBytes_withoutVariable_returnsError(t *testing.T) {
	_, err := NewBytesReferenceBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
