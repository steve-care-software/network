package links

import (
	"reflect"
	"testing"
)

func TestCondition_Success(t *testing.T) {
	resource := NewConditionResourceForTests(23)
	operator := NewOperatorWithAndForTests()
	next := NewConditionValueWithResourceForTests(
		NewConditionResourceForTests(44),
	)

	condition := NewConditionForTests(
		resource,
		operator,
		next,
	)

	retResource := condition.Resource()
	if !reflect.DeepEqual(resource, retResource) {
		t.Errorf("the resource is invalid")
		return
	}

	retOperator := condition.Operator()
	if !reflect.DeepEqual(operator, retOperator) {
		t.Errorf("the operator is invalid")
		return
	}

	retNext := condition.Next()
	if !reflect.DeepEqual(next, retNext) {
		t.Errorf("the next is invalid")
		return
	}
}

func TestCondition__withoutResource_returnsError(t *testing.T) {
	operator := NewOperatorWithAndForTests()
	next := NewConditionValueWithResourceForTests(
		NewConditionResourceForTests(12),
	)

	_, err := NewConditionBuilder().Create().WithOperator(operator).WithNext(next).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCondition__withoutOperator_returnsError(t *testing.T) {
	resource := NewConditionResourceForTests(32)
	next := NewConditionValueWithResourceForTests(
		NewConditionResourceForTests(67),
	)

	_, err := NewConditionBuilder().Create().WithResource(resource).WithNext(next).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestCondition__withoutNext_returnsError(t *testing.T) {
	resource := NewConditionResourceForTests(21)
	operator := NewOperatorWithAndForTests()
	_, err := NewConditionBuilder().Create().WithResource(resource).WithOperator(operator).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
