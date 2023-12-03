package values

import "errors"

type builder struct {
	isNil      bool
	pIntValue  *int
	pRealValue *float64
	textValue  string
	bytes      []byte
}

func createBuilder() Builder {
	out := builder{
		isNil:      false,
		pIntValue:  nil,
		pRealValue: nil,
		textValue:  "",
		bytes:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithInteger adds an integer to the builder
func (app *builder) WithInteger(intValue int) Builder {
	app.pIntValue = &intValue
	return app
}

// WithReal adds a real to the builder
func (app *builder) WithReal(realValue float64) Builder {
	app.pRealValue = &realValue
	return app
}

// WithText adds a text to the builder
func (app *builder) WithText(text string) Builder {
	app.textValue = text
	return app
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes []byte) Builder {
	app.bytes = bytes
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	if app.isNil {
		return createValueWithNil(), nil
	}

	if app.pIntValue != nil {
		return createValueWithInteger(app.pIntValue), nil
	}

	if app.pRealValue != nil {
		return createValueWithReal(app.pRealValue), nil
	}

	if app.textValue != "" {
		return createValueWithText(app.textValue), nil
	}

	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes != nil {
		return createValueWithBytes(app.bytes), nil
	}

	return nil, errors.New("the Value is invalid")
}
