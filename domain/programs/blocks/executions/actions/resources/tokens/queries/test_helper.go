package queries

import (
	"steve.care/network/domain/queries"
	"steve.care/network/domain/queries/conditions"
)

// NewQueryWithQueryForTests creates a new query with query for tests
func NewQueryWithQueryForTests(input queries.Query) Query {
	ins, err := NewBuilder().Create().WithQuery(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewQueryWithConditionForTests creates a new query with condition for tests
func NewQueryWithConditionForTests(input conditions.Condition) Query {
	ins, err := NewBuilder().Create().WithCondition(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewQueryWithPointerForTests creates a new query with pointer for tests
func NewQueryWithPointerForTests(input conditions.Pointer) Query {
	ins, err := NewBuilder().Create().WithPointer(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewQueryWithElementForTests creates a new query with element for tests
func NewQueryWithElementForTests(input conditions.Element) Query {
	ins, err := NewBuilder().Create().WithElement(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewQueryWithResourceForTests creates a new query with resource for tests
func NewQueryWithResourceForTests(input conditions.Resource) Query {
	ins, err := NewBuilder().Create().WithResource(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewQueryWithOperatorForTests creates a new query with operator for tests
func NewQueryWithOperatorForTests(input conditions.Operator) Query {
	ins, err := NewBuilder().Create().WithOperator(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewQueryWithRelationalOperatorForTests creates a new query with relationalOperator for tests
func NewQueryWithRelationalOperatorForTests(input conditions.RelationalOperator) Query {
	ins, err := NewBuilder().Create().WithRelationalOperator(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewQueryWithIntegerOperatorForTests creates a new query with integerOperator for tests
func NewQueryWithIntegerOperatorForTests(input conditions.IntegerOperator) Query {
	ins, err := NewBuilder().Create().WithIntegerOperator(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
