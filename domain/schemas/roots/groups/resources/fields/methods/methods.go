package methods

type methods struct {
	retriever []string
	builder   string
}

func createMethods(
	retriever []string,
	builder string,
) Methods {
	out := methods{
		retriever: retriever,
		builder:   builder,
	}

	return &out
}

// Retriever returns the retriever method
func (obj *methods) Retriever() []string {
	return obj.retriever
}

// Retriever returns the builder method
func (obj *methods) Builder() string {
	return obj.builder
}
