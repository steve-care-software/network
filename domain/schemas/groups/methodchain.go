package groups

type methodChain struct {
	condition string
	retriever []string
	element   Element
}

func createMethodChain(
	condition string,
	retriever []string,
	element Element,
) MethodChain {
	out := methodChain{
		condition: condition,
		retriever: retriever,
		element:   element,
	}

	return &out
}

// Condition returns the condition
func (obj *methodChain) Condition() string {
	return obj.condition
}

// Retriever returns the retriever
func (obj *methodChain) Retriever() []string {
	return obj.retriever
}

// Element returns the value
func (obj *methodChain) Element() Element {
	return obj.element
}
