package layers

import (
	"reflect"
	"testing"
)

func TestEngine_withExecution_Success(t *testing.T) {
	execution := NewBytesReferenceWithVariableForTests("myVariable")
	ins := NewEngineWithExecutionForTests(execution)

	if !ins.IsExecution() {
		t.Errorf("the engine was expected to contain an execution")
		return
	}

	if ins.IsResource() {
		t.Errorf("the engine was expected to NOT contain a resource")
		return
	}

	retExecution := ins.Execution()
	if !reflect.DeepEqual(execution, retExecution) {
		t.Errorf("the execution is invalid")
		return
	}
}

func TestEngine_withResource_Success(t *testing.T) {
	resource := NewAssignableResourceWithCompileForTests(
		NewBytesReferenceWithVariableForTests("myVariable"),
	)

	ins := NewEngineWithResourceForTests(resource)

	if ins.IsExecution() {
		t.Errorf("the engine was expected to NOT contain an execution")
		return
	}

	if !ins.IsResource() {
		t.Errorf("the engine was expected to contain a resource")
		return
	}

	retResource := ins.Resource()
	if !reflect.DeepEqual(resource, retResource) {
		t.Errorf("the resource is invalid")
		return
	}
}

func TestEngine_withoutParam_returnsError(t *testing.T) {
	_, err := NewEngineBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
