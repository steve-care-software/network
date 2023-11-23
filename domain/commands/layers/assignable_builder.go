package layers

import (
	"errors"

	"steve.care/network/domain/hash"
)

type assignableBuilder struct {
	hashAdapter hash.Adapter
	bytes       Bytes
	identity    Identity
}

func createAssignableBuilder(
	hashAdapter hash.Adapter,
) AssignableBuilder {
	out := assignableBuilder{
		hashAdapter: hashAdapter,
		bytes:       nil,
		identity:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignableBuilder) Create() AssignableBuilder {
	return createAssignableBuilder(
		app.hashAdapter,
	)
}

// WithBytes add bytes to the builder
func (app *assignableBuilder) WithBytes(bytes Bytes) AssignableBuilder {
	app.bytes = bytes
	return app
}

// WithIdentity add identity to the builder
func (app *assignableBuilder) WithIdentity(identity Identity) AssignableBuilder {
	app.identity = identity
	return app
}

// Now builds a new Assignable instance
func (app *assignableBuilder) Now() (Assignable, error) {
	data := [][]byte{}
	if app.bytes != nil {
		data = append(data, []byte("bytes"))
		data = append(data, app.bytes.Hash().Bytes())
	}

	if app.identity != nil {
		data = append(data, []byte("identity"))
		data = append(data, app.identity.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Assignable is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.bytes != nil {
		return createAssignableWithBytes(*pHash, app.bytes), nil
	}

	return createAssignableWithIdentity(*pHash, app.identity), nil
}
