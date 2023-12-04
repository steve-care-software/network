package links

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/links"
)

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
	IsOriginResource() links.OriginResource
	IsOperator() bool
	Operator() links.Operator
}
