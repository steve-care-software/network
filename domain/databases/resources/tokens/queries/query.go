package queries

import (
	"steve.care/network/domain/databases/criterias"
	"steve.care/network/domain/databases/criterias/conditions"
	"steve.care/network/domain/hash"
)

type query struct {
	criteria    criterias.Criteria
	condition   conditions.Condition
	pointer     conditions.Pointer
	element     conditions.Element
	resource    conditions.Resource
	operator    conditions.Operator
	relOperator conditions.RelationalOperator
	intOperator conditions.IntegerOperator
}

func createQueryWithCriteria(
	criteria criterias.Criteria,
) Query {
	return createQueryInternally(
		criteria,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createQueryWithCondition(
	condition conditions.Condition,
) Query {
	return createQueryInternally(
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

func createQueryWithPointer(
	pointer conditions.Pointer,
) Query {
	return createQueryInternally(
		nil,
		nil,
		pointer,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createQueryWithElement(
	element conditions.Element,
) Query {
	return createQueryInternally(
		nil,
		nil,
		nil,
		element,
		nil,
		nil,
		nil,
		nil,
	)
}

func createQueryWithResource(
	resource conditions.Resource,
) Query {
	return createQueryInternally(
		nil,
		nil,
		nil,
		nil,
		resource,
		nil,
		nil,
		nil,
	)
}

func createQueryWithOperator(
	operator conditions.Operator,
) Query {
	return createQueryInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		operator,
		nil,
		nil,
	)
}

func createQueryWithRelationalOperator(
	relOperator conditions.RelationalOperator,
) Query {
	return createQueryInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		relOperator,
		nil,
	)
}

func createQueryWithIntegerOperator(
	intOperator conditions.IntegerOperator,
) Query {
	return createQueryInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		intOperator,
	)
}

func createQueryInternally(
	criteria criterias.Criteria,
	condition conditions.Condition,
	pointer conditions.Pointer,
	element conditions.Element,
	resource conditions.Resource,
	operator conditions.Operator,
	relOperator conditions.RelationalOperator,
	intOperator conditions.IntegerOperator,
) Query {
	out := query{
		criteria:    criteria,
		condition:   condition,
		pointer:     pointer,
		element:     element,
		resource:    resource,
		operator:    operator,
		relOperator: relOperator,
		intOperator: intOperator,
	}

	return &out
}

// Hash returns the hash
func (obj *query) Hash() hash.Hash {
	if obj.IsCriteria() {
		return obj.criteria.Hash()
	}

	if obj.IsCondition() {
		return obj.condition.Hash()
	}

	if obj.IsPointer() {
		return obj.pointer.Hash()
	}

	if obj.IsElement() {
		return obj.element.Hash()
	}

	if obj.IsResource() {
		return obj.resource.Hash()
	}

	if obj.IsOperator() {
		return obj.operator.Hash()
	}

	if obj.IsRelationalOperator() {
		return obj.RelationalOperator().Hash()
	}

	return obj.intOperator.Hash()
}

// IsCriteria returns true if there is a criteria, false otherwise
func (obj *query) IsCriteria() bool {
	return obj.criteria != nil
}

// Criteria returns the criteria, if any
func (obj *query) Criteria() criterias.Criteria {
	return obj.criteria
}

// IsCondition returns true if there is a condition, false otherwise
func (obj *query) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *query) Condition() conditions.Condition {
	return obj.condition
}

// IsPointer returns true if there is a pointer, false otherwise
func (obj *query) IsPointer() bool {
	return obj.pointer != nil
}

// Pointer returns the pointer, if any
func (obj *query) Pointer() conditions.Pointer {
	return obj.pointer
}

// IsElement returns true if there is an element, false otherwise
func (obj *query) IsElement() bool {
	return obj.element != nil
}

// Element returns the element, if any
func (obj *query) Element() conditions.Element {
	return obj.element
}

// IsResource returns true if there is a resource, false otherwise
func (obj *query) IsResource() bool {
	return obj.resource != nil
}

// Resource returns the resource, if any
func (obj *query) Resource() conditions.Resource {
	return obj.resource
}

// IsOperator returns true if there is a operator, false otherwise
func (obj *query) IsOperator() bool {
	return obj.operator != nil
}

// Operator returns the operator, if any
func (obj *query) Operator() conditions.Operator {
	return obj.operator
}

// IsRelationalOperator returns true if there is a relational operator, false otherwise
func (obj *query) IsRelationalOperator() bool {
	return obj.relOperator != nil
}

// RelationalOperator returns the relational operator, if any
func (obj *query) RelationalOperator() conditions.RelationalOperator {
	return obj.relOperator
}

// IsIntegerOperator returns true if there is an integer operator, false otherwise
func (obj *query) IsIntegerOperator() bool {
	return obj.intOperator != nil
}

// IntegerOperator returns the integer operator, if any
func (obj *query) IntegerOperator() conditions.IntegerOperator {
	return obj.intOperator
}
