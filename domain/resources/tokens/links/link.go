package links

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/links"
)

type link struct {
	link              links.Link
	element           links.Element
	condition         links.Condition
	conditionValue    links.ConditionValue
	conditionResource links.ConditionResource
	origin            links.Origin
	originValue       links.OriginValue
	originResource    links.OriginResource
	operator          links.Operator
}

func createLinkWithLink(
	linkIns links.Link,
) Link {
	return createLinkInternally(
		linkIns,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLinkWithElement(
	element links.Element,
) Link {
	return createLinkInternally(
		nil,
		element,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLinkWithCondition(
	condition links.Condition,
) Link {
	return createLinkInternally(
		nil,
		nil,
		condition,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLinkWithConditionValue(
	conditionValue links.ConditionValue,
) Link {
	return createLinkInternally(
		nil,
		nil,
		nil,
		conditionValue,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLinkWithConditionResource(
	conditionResource links.ConditionResource,
) Link {
	return createLinkInternally(
		nil,
		nil,
		nil,
		nil,
		conditionResource,
		nil,
		nil,
		nil,
		nil,
	)
}

func createLinkWithOrigin(
	origin links.Origin,
) Link {
	return createLinkInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		origin,
		nil,
		nil,
		nil,
	)
}

func createLinkWithOriginValue(
	originValue links.OriginValue,
) Link {
	return createLinkInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		originValue,
		nil,
		nil,
	)
}

func createLinkWithOriginResource(
	originResource links.OriginResource,
) Link {
	return createLinkInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		originResource,
		nil,
	)
}

func createLinkWithOperator(
	operator links.Operator,
) Link {
	return createLinkInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		operator,
	)
}

func createLinkInternally(
	linkIns links.Link,
	element links.Element,
	condition links.Condition,
	conditionValue links.ConditionValue,
	conditionResource links.ConditionResource,
	origin links.Origin,
	originValue links.OriginValue,
	originResource links.OriginResource,
	operator links.Operator,
) Link {
	out := link{
		link:              linkIns,
		element:           element,
		condition:         condition,
		conditionValue:    conditionValue,
		conditionResource: conditionResource,
		origin:            origin,
		originValue:       originValue,
		originResource:    originResource,
		operator:          operator,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	if obj.IsLink() {
		return obj.link.Hash()
	}

	if obj.IsElement() {
		return obj.element.Hash()
	}

	if obj.IsCondition() {
		return obj.condition.Hash()
	}

	if obj.IsConditionValue() {
		return obj.conditionValue.Hash()
	}

	if obj.IsConditionResource() {
		return obj.conditionResource.Hash()
	}

	if obj.IsOrigin() {
		return obj.origin.Hash()
	}

	if obj.IsOriginValue() {
		return obj.originValue.Hash()
	}

	if obj.IsOriginResource() {
		return obj.originResource.Hash()
	}

	return obj.operator.Hash()
}

// IsLink returns true if there is a link, false otherwise
func (obj *link) IsLink() bool {
	return obj.link != nil
}

// Link returns the link, if any
func (obj *link) Link() links.Link {
	return obj.link
}

// IsElement returns true if there is an element, false otherwise
func (obj *link) IsElement() bool {
	return obj.element != nil
}

// Element returns the element, if any
func (obj *link) Element() links.Element {
	return obj.element
}

// IsCondition returns true if there is a condition, false otherwise
func (obj *link) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *link) Condition() links.Condition {
	return obj.condition
}

// IsConditionValue returns true if there is a conditionValue, false otherwise
func (obj *link) IsConditionValue() bool {
	return obj.conditionValue != nil
}

// ConditionValue returns the conditionValue, if any
func (obj *link) ConditionValue() links.ConditionValue {
	return obj.conditionValue
}

// IsConditionResource returns true if there is a conditionResource, false otherwise
func (obj *link) IsConditionResource() bool {
	return obj.conditionResource != nil
}

// ConditionResource returns the conditionResource, if any
func (obj *link) ConditionResource() links.ConditionResource {
	return obj.conditionResource
}

// IsOrigin returns true if there is an origin, false otherwise
func (obj *link) IsOrigin() bool {
	return obj.origin != nil
}

// Origin returns the origin, if any
func (obj *link) Origin() links.Origin {
	return obj.origin
}

// IsOriginValue returns true if there is an originValue, false otherwise
func (obj *link) IsOriginValue() bool {
	return obj.originValue != nil
}

// OriginValue returns the originValue, if any
func (obj *link) OriginValue() links.OriginValue {
	return obj.originValue
}

// IsOriginResource returns true if there is an originResource, false otherwise
func (obj *link) IsOriginResource() bool {
	return obj.originResource != nil
}

// OriginResource returns the originResource, if any
func (obj *link) OriginResource() links.OriginResource {
	return obj.originResource
}

// IsOperator returns true if there is an operator, false otherwise
func (obj *link) IsOperator() bool {
	return obj.operator != nil
}

// Operator returns the operator, if any
func (obj *link) Operator() links.Operator {
	return obj.operator
}
