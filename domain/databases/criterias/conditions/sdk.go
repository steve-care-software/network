package conditions

import "steve.care/network/domain/databases/criterias/values"

// NewRelationalOperatorBuilder creates a new relational operator builder
func NewRelationalOperatorBuilder() RelationalOperatorBuilder {
	return createRelationalOperatorBuilder()
}

// NewIntegerOperatorBuilder creates a new integer operator builder
func NewIntegerOperatorBuilder() IntegerOperatorBuilder {
	return createIntegerOperatorBuilder()
}

// Builder represents a condition builder
type Builder interface {
	Create() Builder
	WithPointer(pointer Pointer) Builder
	WithOperator(operator Operator) Builder
	WithElement(element Element) Builder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Pointer() Pointer
	Operator() Operator
	Element() Element
}

// PointerBuilder represents a pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithContainer(container string) PointerBuilder
	WithField(field string) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a field pointer
type Pointer interface {
	Container() string
	Field() string
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithCondition(condition Condition) ElementBuilder
	WithResource(resource Resource) ElementBuilder
	Now() (Element, error)
}

// Element represents a conditional element
type Element interface {
	IsCondition() bool
	Condition() Condition
	IsResource() bool
	Resource() Resource
}

// ResourceBuilder represents a resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithField(field Pointer) ResourceBuilder
	WithValue(value values.Value) ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	IsField() bool
	Field() Pointer
	IsValue() bool
	Value() values.Value
}

// OperatorBuilder represents an operator builder
type OperatorBuilder interface {
	Create() OperatorBuilder
	WithRelational(relational RelationalOperator) OperatorBuilder
	WithInteger(integer IntegerOperator) OperatorBuilder
	IsEqual() OperatorBuilder
	Now() (Operator, error)
}

// Operator represents an operator
type Operator interface {
	IsEqual() bool
	IsRelational() bool
	Relational() RelationalOperator
	IsInteger() bool
	Integer() IntegerOperator
}

// RelationalOperatorBuilder represents a relational operator builder
type RelationalOperatorBuilder interface {
	Create() RelationalOperatorBuilder
	IsAnd() RelationalOperatorBuilder
	IsOr() RelationalOperatorBuilder
	Now() (RelationalOperator, error)
}

// RelationalOperator represents a relational operator
type RelationalOperator interface {
	IsAnd() bool
	IsOr() bool
}

// IntegerOperatorBuilder represents the integer operator builder
type IntegerOperatorBuilder interface {
	Create() IntegerOperatorBuilder
	IsSmallerThan() IntegerOperatorBuilder
	IsBiggerThan() IntegerOperatorBuilder
	IsEqual() IntegerOperatorBuilder
	Now() (IntegerOperator, error)
}

// IntegerOperator represents an integer operator
type IntegerOperator interface {
	IsSmallerThan() bool
	IsBiggerThan() bool
	IsEqual() bool
}
