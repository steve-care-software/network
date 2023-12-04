package kinds

import "errors"

type builder struct {
	isNil     bool
	isInteger bool
	isReal    bool
	isText    bool
	isBlob    bool
}

func createBuilder() Builder {
	out := builder{
		isNil:     false,
		isInteger: false,
		isReal:    false,
		isText:    false,
		isBlob:    false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// IsNil flags the builder as nil
func (app *builder) IsNil() Builder {
	app.isNil = true
	return app
}

// IsInteger flags the builder as integer
func (app *builder) IsInteger() Builder {
	app.isInteger = true
	return app
}

// IsReal flags the builder as real
func (app *builder) IsReal() Builder {
	app.isReal = true
	return app
}

// IsText flags the builder as text
func (app *builder) IsText() Builder {
	app.isText = true
	return app
}

// IsBlob flags the builder as blob
func (app *builder) IsBlob() Builder {
	app.isBlob = true
	return app
}

// Now builds a new Kind instance
func (app *builder) Now() (Kind, error) {
	if app.isNil {
		return createKindWithNil(), nil
	}

	if app.isInteger {
		return createKindWithInteger(), nil
	}

	if app.isReal {
		return createKindWithReal(), nil
	}

	if app.isText {
		return createKindWithText(), nil
	}

	if app.isBlob {
		return createKindWithBlob(), nil
	}

	return nil, errors.New("the Kind is invalid")
}
