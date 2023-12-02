package conditions

import "steve.care/network/domain/databases/values"

// Condition represents a condition
type Condition interface {
	Pointer() Pointer
	Operator() Operator
	Element() Element
}

// Pointer represents a field pointer
type Pointer interface {
	Container() string
	Field() string
}

// Element represents a conditional element
type Element interface {
	IsCondition() bool
	Condition() Condition
	IsResource() bool
	Resource() Resource
}

// Resource represents a resource
type Resource interface {
	IsField() bool
	Field() Pointer
	IsValue() bool
	Value() values.Value
}

// Operator represents an operator
type Operator interface {
	IsEqual() bool
	IsRelational() bool
	Relational() RelationalOperator
	IsInteger() bool
	Integer() IntegerOperator
}

// RelationalOperator represents a relational operator
type RelationalOperator interface {
	IsAnd() bool
	IsOr() bool
}

// IntegerOperator represents an integer operator
type IntegerOperator interface {
	IsSmallerThan() bool
	IsBiggerThan() bool
	HasEqual() bool
}
