package conditions

import "errors"

type builder struct {
	pointer  Pointer
	operator Operator
	element  Element
}

func createBuilder() Builder {
	out := builder{
		pointer:  nil,
		operator: nil,
		element:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithPointer adds a pointer to the builder
func (app *builder) WithPointer(pointer Pointer) Builder {
	app.pointer = pointer
	return app
}

// WithOperator adds an operator to the builder
func (app *builder) WithOperator(operator Operator) Builder {
	app.operator = operator
	return app
}

// WithElement adds an element to the builder
func (app *builder) WithElement(element Element) Builder {
	app.element = element
	return app
}

// Now builds a new Condition instance
func (app *builder) Now() (Condition, error) {
	if app.pointer == nil {
		return nil, errors.New("the pointer is mandatory in order to build a Condition instance")
	}

	if app.operator == nil {
		return nil, errors.New("the operator is mandatory in order to build a Condition instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Condition instance")
	}

	return createCondition(
		app.pointer,
		app.operator,
		app.element,
	), nil
}
