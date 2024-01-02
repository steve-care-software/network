package links

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers/links"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the link builder
type Builder interface {
	Create() Builder
	WithLink(link links.Link) Builder
	WithElement(elemnt links.Element) Builder
	WithCondition(condition links.Condition) Builder
	WithConditionValue(conditionValue links.ConditionValue) Builder
	WithConditionResource(conditionResource links.ConditionResource) Builder
	WithOrigin(origin links.Origin) Builder
	WithOriginValue(originValue links.OriginValue) Builder
	WithOriginResource(originResource links.OriginResource) Builder
	WithOperator(operator links.Operator) Builder
	Now() (Link, error)
}

// Link represents a link resource
type Link interface {
	Hash() hash.Hash
	IsLink() bool
	Link() links.Link
	IsElement() bool
	Element() links.Element
	IsCondition() bool
	Condition() links.Condition
	IsConditionValue() bool
	ConditionValue() links.ConditionValue
	IsConditionResource() bool
	ConditionResource() links.ConditionResource
	IsOrigin() bool
	Origin() links.Origin
	IsOriginValue() bool
	OriginValue() links.OriginValue
	IsOriginResource() bool
	OriginResource() links.OriginResource
	IsOperator() bool
	Operator() links.Operator
}
