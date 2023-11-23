package layers

import (
	"errors"

	"steve.care/network/domain/hash"
)

type bytesBuilder struct {
	hashAdapter hash.Adapter
	join        BytesReferences
	compare     BytesReferences
}

func createBytesBuilder(
	hashAdapter hash.Adapter,
) BytesBuilder {
	out := bytesBuilder{
		hashAdapter: hashAdapter,
		join:        nil,
		compare:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *bytesBuilder) Create() BytesBuilder {
	return createBytesBuilder(
		app.hashAdapter,
	)
}

// WithJoin adds a join to the builder
func (app *bytesBuilder) WithJoin(join BytesReferences) BytesBuilder {
	app.join = join
	return app
}

// WithCompare adds a compare to the builder
func (app *bytesBuilder) WithCompare(compare BytesReferences) BytesBuilder {
	app.compare = compare
	return app
}

// Now builds a new Bytes instance
func (app *bytesBuilder) Now() (Bytes, error) {
	data := [][]byte{}
	if app.join != nil {
		data = append(data, []byte("join"))
		data = append(data, app.join.Hash().Bytes())
	}

	if app.compare != nil {
		data = append(data, []byte("compare"))
		data = append(data, app.compare.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Bytes is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.join != nil {
		return createBytesWithJoin(*pHash, app.join), nil
	}

	return createBytesWithCompare(*pHash, app.compare), nil
}
