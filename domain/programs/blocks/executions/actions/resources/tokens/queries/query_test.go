package queries

import (
	"reflect"
	"testing"

	"steve.care/network/domain/queries"
	"steve.care/network/domain/queries/conditions"
)

func TestQuery_withQuery_Success(t *testing.T) {
	query := queries.NewQueryForTests(
		"myEntity",
		conditions.NewConditionForTests(
			conditions.NewPointerForTests("myEntity", "myField"),
			conditions.NewOperatorWithEqualForTests(),
			conditions.NewElementWithResourceForTests(
				conditions.NewResourceWithValueForTests(45),
			),
		),
	)

	ins := NewQueryWithQueryForTests(query)

	if !ins.IsQuery() {
		t.Errorf("the query was expected to contain a query")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the query was expected to NOT contain a condition")
		return
	}

	if ins.IsPointer() {
		t.Errorf("the query was expected to NOT contain a pointer")
		return
	}

	if ins.IsElement() {
		t.Errorf("the query was expected to NOT contain an element")
		return
	}

	if ins.IsResource() {
		t.Errorf("the query was expected to NOT contain a resource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the query was expected to NOT contain an operator")
		return
	}

	if ins.IsRelationalOperator() {
		t.Errorf("the query was expected to NOT contain a relationalOperator")
		return
	}

	if ins.IsIntegerOperator() {
		t.Errorf("the query was expected to NOT contain an integerOperator")
		return
	}

	retQuery := ins.Query()
	if !reflect.DeepEqual(query, retQuery) {
		t.Errorf("the returned query is invalid")
		return
	}
}

func TestQuery_withCondition_Success(t *testing.T) {
	condition := conditions.NewConditionForTests(
		conditions.NewPointerForTests("myEntity", "myField"),
		conditions.NewOperatorWithEqualForTests(),
		conditions.NewElementWithResourceForTests(
			conditions.NewResourceWithValueForTests(45),
		),
	)

	ins := NewQueryWithConditionForTests(condition)

	if ins.IsQuery() {
		t.Errorf("the query was expected to NOT contain a query")
		return
	}

	if !ins.IsCondition() {
		t.Errorf("the query was expected to contain a condition")
		return
	}

	if ins.IsPointer() {
		t.Errorf("the query was expected to NOT contain a pointer")
		return
	}

	if ins.IsElement() {
		t.Errorf("the query was expected to NOT contain an element")
		return
	}

	if ins.IsResource() {
		t.Errorf("the query was expected to NOT contain a resource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the query was expected to NOT contain an operator")
		return
	}

	if ins.IsRelationalOperator() {
		t.Errorf("the query was expected to NOT contain a relationalOperator")
		return
	}

	if ins.IsIntegerOperator() {
		t.Errorf("the query was expected to NOT contain an integerOperator")
		return
	}

	retCondition := ins.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the returned condition is invalid")
		return
	}
}

func TestQuery_withPointer_Success(t *testing.T) {
	pointer := conditions.NewPointerForTests("myEntity", "myField")
	ins := NewQueryWithPointerForTests(pointer)

	if ins.IsQuery() {
		t.Errorf("the query was expected to NOT contain a query")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the query was expected to NOT contain a condition")
		return
	}

	if !ins.IsPointer() {
		t.Errorf("the query was expected to contain a pointer")
		return
	}

	if ins.IsElement() {
		t.Errorf("the query was expected to NOT contain an element")
		return
	}

	if ins.IsResource() {
		t.Errorf("the query was expected to NOT contain a resource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the query was expected to NOT contain an operator")
		return
	}

	if ins.IsRelationalOperator() {
		t.Errorf("the query was expected to NOT contain a relationalOperator")
		return
	}

	if ins.IsIntegerOperator() {
		t.Errorf("the query was expected to NOT contain an integerOperator")
		return
	}

	retPointer := ins.Pointer()
	if !reflect.DeepEqual(pointer, retPointer) {
		t.Errorf("the returned pointer is invalid")
		return
	}
}

func TestQuery_withElement_Success(t *testing.T) {
	element := conditions.NewElementWithResourceForTests(
		conditions.NewResourceWithFieldForTests(
			conditions.NewPointerForTests("myEntity", "myField"),
		),
	)

	ins := NewQueryWithElementForTests(element)

	if ins.IsQuery() {
		t.Errorf("the query was expected to NOT contain a query")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the query was expected to NOT contain a condition")
		return
	}

	if ins.IsPointer() {
		t.Errorf("the query was expected to NOT contain a pointer")
		return
	}

	if !ins.IsElement() {
		t.Errorf("the query was expected to contain an element")
		return
	}

	if ins.IsResource() {
		t.Errorf("the query was expected to NOT contain a resource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the query was expected to NOT contain an operator")
		return
	}

	if ins.IsRelationalOperator() {
		t.Errorf("the query was expected to NOT contain a relationalOperator")
		return
	}

	if ins.IsIntegerOperator() {
		t.Errorf("the query was expected to NOT contain an integerOperator")
		return
	}

	retElement := ins.Element()
	if !reflect.DeepEqual(element, retElement) {
		t.Errorf("the returned element is invalid")
		return
	}
}

func TestQuery_withResource_Success(t *testing.T) {
	resource := conditions.NewResourceWithFieldForTests(
		conditions.NewPointerForTests("myEntity", "myField"),
	)

	ins := NewQueryWithResourceForTests(resource)

	if ins.IsQuery() {
		t.Errorf("the query was expected to NOT contain a query")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the query was expected to NOT contain a condition")
		return
	}

	if ins.IsPointer() {
		t.Errorf("the query was expected to NOT contain a pointer")
		return
	}

	if ins.IsElement() {
		t.Errorf("the query was expected to NOT contain an element")
		return
	}

	if !ins.IsResource() {
		t.Errorf("the query was expected to contain a resource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the query was expected to NOT contain an operator")
		return
	}

	if ins.IsRelationalOperator() {
		t.Errorf("the query was expected to NOT contain a relationalOperator")
		return
	}

	if ins.IsIntegerOperator() {
		t.Errorf("the query was expected to NOT contain an integerOperator")
		return
	}

	retResource := ins.Resource()
	if !reflect.DeepEqual(resource, retResource) {
		t.Errorf("the returned resource is invalid")
		return
	}
}

func TestQuery_withOperator_Success(t *testing.T) {
	operator := conditions.NewOperatorWithEqualForTests()
	ins := NewQueryWithOperatorForTests(operator)

	if ins.IsQuery() {
		t.Errorf("the query was expected to NOT contain a query")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the query was expected to NOT contain a condition")
		return
	}

	if ins.IsPointer() {
		t.Errorf("the query was expected to NOT contain a pointer")
		return
	}

	if ins.IsElement() {
		t.Errorf("the query was expected to NOT contain an element")
		return
	}

	if ins.IsResource() {
		t.Errorf("the query was expected to NOT contain a resource")
		return
	}

	if !ins.IsOperator() {
		t.Errorf("the query was expected to contain an operator")
		return
	}

	if ins.IsRelationalOperator() {
		t.Errorf("the query was expected to NOT contain a relationalOperator")
		return
	}

	if ins.IsIntegerOperator() {
		t.Errorf("the query was expected to NOT contain an integerOperator")
		return
	}

	retOperator := ins.Operator()
	if !reflect.DeepEqual(operator, retOperator) {
		t.Errorf("the returned operator is invalid")
		return
	}
}

func TestQuery_withRelationalOperator_Success(t *testing.T) {
	relationalOperator := conditions.NewRelationalOperatorWithAndForTests()
	ins := NewQueryWithRelationalOperatorForTests(relationalOperator)

	if ins.IsQuery() {
		t.Errorf("the query was expected to NOT contain a query")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the query was expected to NOT contain a condition")
		return
	}

	if ins.IsPointer() {
		t.Errorf("the query was expected to NOT contain a pointer")
		return
	}

	if ins.IsElement() {
		t.Errorf("the query was expected to NOT contain an element")
		return
	}

	if ins.IsResource() {
		t.Errorf("the query was expected to NOT contain a resource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the query was expected to NOT contain an operator")
		return
	}

	if !ins.IsRelationalOperator() {
		t.Errorf("the query was expected to contain a relationalOperator")
		return
	}

	if ins.IsIntegerOperator() {
		t.Errorf("the query was expected to NOT contain an integerOperator")
		return
	}

	retRelationalOperator := ins.RelationalOperator()
	if !reflect.DeepEqual(relationalOperator, retRelationalOperator) {
		t.Errorf("the returned relationalOperator is invalid")
		return
	}
}

func TestQuery_withIntegerOperator_Success(t *testing.T) {
	integerOperator := conditions.NewIntegerOperatorWithIsSmallerThanAndIsEqualForTests()
	ins := NewQueryWithIntegerOperatorForTests(integerOperator)

	if ins.IsQuery() {
		t.Errorf("the query was expected to NOT contain a query")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the query was expected to NOT contain a condition")
		return
	}

	if ins.IsPointer() {
		t.Errorf("the query was expected to NOT contain a pointer")
		return
	}

	if ins.IsElement() {
		t.Errorf("the query was expected to NOT contain an element")
		return
	}

	if ins.IsResource() {
		t.Errorf("the query was expected to NOT contain a resource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the query was expected to NOT contain an operator")
		return
	}

	if ins.IsRelationalOperator() {
		t.Errorf("the query was expected to NOT contain a relationalOperator")
		return
	}

	if !ins.IsIntegerOperator() {
		t.Errorf("the query was expected to contain an integerOperator")
		return
	}

	retIntegerOperator := ins.IntegerOperator()
	if !reflect.DeepEqual(integerOperator, retIntegerOperator) {
		t.Errorf("the returned integerOperator is invalid")
		return
	}
}

func TestQuery_withoutParam_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
