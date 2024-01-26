package groups

import "errors"

type methodChainBuilder struct {
	condition string
	value     string
	element   Element
}

func createMethodChainBuilder() MethodChainBuilder {
	out := methodChainBuilder{
		condition: "",
		value:     "",
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

// WithValue adds a value to the builder
func (app *methodChainBuilder) WithValue(value string) MethodChainBuilder {
	app.value = value
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

	if app.value == "" {
		return nil, errors.New("the value is mandatory in order to build a MethodChain instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a MethodChain instance")
	}

	return createMethodChain(
		app.condition,
		app.value,
		app.element,
	), nil
}
