package methods

type methods struct {
	retriever []string
	element   string
}

func createMethods(
	retriever []string,
	element string,
) Methods {
	out := methods{
		retriever: retriever,
		element:   element,
	}

	return &out
}

// Retriever returns the retriever method
func (obj *methods) Retriever() []string {
	return obj.retriever
}

// Retriever returns the element method
func (obj *methods) Element() string {
	return obj.element
}
