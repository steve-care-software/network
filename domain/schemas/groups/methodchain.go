package groups

type methodChain struct {
	condition string
	value     string
	element   Element
}

func createMethodChain(
	condition string,
	value string,
	element Element,
) MethodChain {
	out := methodChain{
		condition: condition,
		value:     value,
		element:   element,
	}

	return &out
}

// Condition returns the condition
func (obj *methodChain) Condition() string {
	return obj.condition
}

// Value returns the value
func (obj *methodChain) Value() string {
	return obj.value
}

// Element returns the value
func (obj *methodChain) Element() Element {
	return obj.element
}
