package layers

import (
	"reflect"
	"testing"
)

func TestAssignable_withBytes_Success(t *testing.T) {
	bytes := NewBytesWithJoinForTests(
		NewBytesReferencesForTests(
			[]BytesReference{
				NewBytesReferenceWithVariableForTests("myVariable"),
				NewBytesReferenceWithBytesForTests([]byte("this is some bytes")),
			},
		),
	)

	ins := NewAssignableWithBytesForTests(bytes)

	if !ins.IsBytes() {
		t.Errorf("the bytes was expected to contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the bytes was expected to NOT contain an identity")
		return
	}

	retBytes := ins.Bytes()
	if !reflect.DeepEqual(bytes, retBytes) {
		t.Errorf("the returned bytes is invalid")
		return
	}
}

func TestAssignable_withIdentity_Success(t *testing.T) {
	sign := NewBytesReferenceWithVariableForTests("mySign")
	signer := NewSignerWithSignForTests(sign)
	identity := NewIdentityWithSignerForTests(signer)

	ins := NewAssignableWithIdentityForTests(identity)

	if ins.IsBytes() {
		t.Errorf("the bytes was expected to NOT contain a bytes")
		return
	}

	if !ins.IsIdentity() {
		t.Errorf("the bytes was expected to contain an identity")
		return
	}

	retIdentity := ins.Identity()
	if !reflect.DeepEqual(identity, retIdentity) {
		t.Errorf("the returned identity is invalid")
		return
	}
}

func TestAssignable_withoutParam_returnsError(t *testing.T) {
	_, err := NewAssignableBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}
