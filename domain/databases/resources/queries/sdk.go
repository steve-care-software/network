package queries

import (
	"steve.care/network/domain/databases/criterias"
	"steve.care/network/domain/databases/criterias/conditions"
	"steve.care/network/domain/hash"
)

// Query represents the query resource
type Query interface {
	Hash() hash.Hash
	IsCriteria() bool
	Criteria() criterias.Criteria
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
