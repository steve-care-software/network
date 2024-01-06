package queries

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/queries"
	"steve.care/network/domain/queries/conditions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a query builder
type Builder interface {
	Create() Builder
	WithQuery(query queries.Query) Builder
	WithCondition(condition conditions.Condition) Builder
	WithPointer(pointer conditions.Pointer) Builder
	WithElement(element conditions.Element) Builder
	WithResource(resource conditions.Resource) Builder
	WithOperator(operator conditions.Operator) Builder
	WithRelationalOperator(relOperator conditions.RelationalOperator) Builder
	WithIntegerOperator(intOperator conditions.IntegerOperator) Builder
	Now() (Query, error)
}

// Query represents the query resource
type Query interface {
	Hash() hash.Hash
	IsQuery() bool
	Query() queries.Query
	IsCondition() bool
	Condition() conditions.Condition
	IsPointer() bool
	Pointer() conditions.Pointer
	IsElement() bool
	Element() conditions.Element
	IsResource() bool
	Resource() conditions.Resource
	IsOperator() bool
	Operator() conditions.Operator
	IsRelationalOperator() bool
	RelationalOperator() conditions.RelationalOperator
	IsIntegerOperator() bool
	IntegerOperator() conditions.IntegerOperator
}
