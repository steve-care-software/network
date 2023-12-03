package conditions

type condition struct {
	pointer  Pointer
	operator Operator
	element  Element
}

func createCondition(
	pointer Pointer,
	operator Operator,
	element Element,
) Condition {
	out := condition{
		pointer:  pointer,
		operator: operator,
		element:  element,
	}

	return &out
}

// Pointer returns the pointer
func (obj *condition) Pointer() Pointer {
	return obj.pointer
}

// Operator returns the operator
func (obj *condition) Operator() Operator {
	return obj.operator
}

// Element returns the element
func (obj *condition) Element() Element {
	return obj.element
}
