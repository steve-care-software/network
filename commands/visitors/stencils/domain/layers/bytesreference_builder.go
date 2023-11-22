package layers

import (
	"errors"

	"steve.care/network/libraries/hash"
)

type bytesReferenceBuilder struct {
	hashAdapter hash.Adapter
	variable    string
	bytes       []byte
}

func createBytesReferenceBuilder(
	hashAdapter hash.Adapter,
) BytesReferenceBuilder {
	out := bytesReferenceBuilder{
		hashAdapter: hashAdapter,
		variable:    "",
		bytes:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *bytesReferenceBuilder) Create() BytesReferenceBuilder {
	return createBytesReferenceBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *bytesReferenceBuilder) WithVariable(variable string) BytesReferenceBuilder {
	app.variable = variable
	return app
}

// WithBytes add bytes to the builder
func (app *bytesReferenceBuilder) WithBytes(bytes []byte) BytesReferenceBuilder {
	app.bytes = bytes
	return app
}

// Now builds a new BytesReference instance
func (app *bytesReferenceBuilder) Now() (BytesReference, error) {
	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	data := []byte{}
	if app.variable != "" {
		data = append(data, []byte(app.variable)...)
	}

	if app.bytes != nil {
		data = append(data, app.bytes...)
	}

	if len(data) <= 0 {
		return nil, errors.New("the BytesReference is invalid")
	}

	pHash, err := app.hashAdapter.FromBytes(data)
	if err != nil {
		return nil, err
	}

	if app.variable != "" {
		return createBytesReferenceWithVariable(*pHash, app.variable), nil
	}

	return createBytesReferenceWithBytes(*pHash, app.bytes), nil
}
