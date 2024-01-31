package methods

import "errors"

type builder struct {
	retriever []string
	builder   string
}

func createBuilder() Builder {
	out := builder{
		retriever: nil,
		builder:   "",
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

// WithBuilder adds a builder to the builder
func (app *builder) WithBuilder(builder string) Builder {
	app.builder = builder
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

	if app.builder == "" {
		return nil, errors.New("the builder is mandatory in order to build a Methods instance")
	}

	return createMethods(app.retriever, app.builder), nil
}
