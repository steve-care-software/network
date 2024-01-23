package kinds

import "errors"

type builder struct {
	isNil   bool
	pInt    *int
	pFloat  *float64
	pString *string
	bytes   []byte
}

func createBuilder() Builder {
	out := builder{
		isNil:   false,
		pInt:    nil,
		pFloat:  nil,
		pString: nil,
		bytes:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithInteger adds an integer to the builder
func (app *builder) WithInteger(intValue int) Builder {
	app.pInt = &intValue
	return app
}

// WithFloat adds a float to the builder
func (app *builder) WithFloat(floatValue float64) Builder {
	app.pFloat = &floatValue
	return app
}

// WithString adds a string to the builder
func (app *builder) WithString(stringValue string) Builder {
	app.pString = &stringValue
	return app
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes []byte) Builder {
	app.bytes = bytes
	return app
}

// IsNil flags the builder as nil
func (app *builder) IsNil() Builder {
	app.isNil = true
	return app
}

// Now builds a new Kind instance
func (app *builder) Now() (Kind, error) {
	if app.isNil {
		return createKindWithNil(), nil
	}

	if app.pInt != nil {
		return createKindWithInteger(app.pInt), nil
	}

	if app.pFloat != nil {
		return createKindWithFloat(app.pFloat), nil
	}

	if app.pString != nil {
		return createKindWithString(app.pString), nil
	}

	if app.bytes != nil {
		return createKindWithBytes(app.bytes), nil
	}

	return nil, errors.New("the Kind is invalid")
}
