package groups

import (
	"errors"
)

type methodChainBuilder struct {
	condition string
	retriever []string
	element   Element
}

func createMethodChainBuilder() MethodChainBuilder {
	out := methodChainBuilder{
		condition: "",
		retriever: nil,
		element:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *methodChainBuilder) Create() MethodChainBuilder {
	return createMethodChainBuilder()
}

// WithCondition adds a condition to the builder
func (app *methodChainBuilder) WithCondition(condition string) MethodChainBuilder {
	app.condition = condition
	return app
}

// WithRetriever adds a retriever to the builder
func (app *methodChainBuilder) WithRetriever(retriever []string) MethodChainBuilder {
	app.retriever = retriever
	return app
}

// WithElement adds an element to the builder
func (app *methodChainBuilder) WithElement(element Element) MethodChainBuilder {
	app.element = element
	return app
}

// Now builds a new MethodChain instance
func (app *methodChainBuilder) Now() (MethodChain, error) {
	if app.condition == "" {
		return nil, errors.New("the condition is mandatory in order to build a MethodChain instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a MethodChain instance")
	}

	if app.retriever != nil && len(app.retriever) <= 0 {
		app.retriever = nil
	}

	if app.retriever == nil {
		return nil, errors.New("the retriever is mandatory in order to build a MethodChain instance")
	}

	return createMethodChain(
		app.condition,
		app.retriever,
		app.element,
	), nil
}
