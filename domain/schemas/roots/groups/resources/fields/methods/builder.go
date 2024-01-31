package methods

import "errors"

type builder struct {
	retriever []string
	element   string
}

func createBuilder() Builder {
	out := builder{
		retriever: nil,
		element:   "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRetriever adds a retriever to the builder
func (app *builder) WithRetriever(retriever []string) Builder {
	app.retriever = retriever
	return app
}

// WithElement adds an element to the builder
func (app *builder) WithElement(element string) Builder {
	app.element = element
	return app
}

// Now builds a new Methods instance
func (app *builder) Now() (Methods, error) {
	if app.retriever != nil && len(app.retriever) <= 0 {
		app.retriever = nil
	}

	if app.retriever == nil {
		return nil, errors.New("the retriever is mandatory in order to build a Methods instance")
	}

	if app.element == "" {
		return nil, errors.New("the element method is mandatory in order to build a Methods instance")
	}

	return createMethods(app.retriever, app.element), nil
}
