package layers

import (
	"errors"

	"steve.care/network/domain/hash"
)

type assignableResourceBuilder struct {
	hashAdapter     hash.Adapter
	compile         BytesReference
	decompile       string
	amountByQuery   BytesReference
	retrieveByQuery BytesReference
	retrieveByHash  BytesReference
	isAmount        bool
}

func createAssignableResourceBuilder(
	hashAdapter hash.Adapter,
) AssignableResourceBuilder {
	out := assignableResourceBuilder{
		hashAdapter:     hashAdapter,
		compile:         nil,
		decompile:       "",
		amountByQuery:   nil,
		retrieveByQuery: nil,
		retrieveByHash:  nil,
		isAmount:        false,
	}

	return &out
}

// Create initializes the builder
func (app *assignableResourceBuilder) Create() AssignableResourceBuilder {
	return createAssignableResourceBuilder(
		app.hashAdapter,
	)
}

// WithCompile adds a compile to the builder
func (app *assignableResourceBuilder) WithCompile(compile BytesReference) AssignableResourceBuilder {
	app.compile = compile
	return app
}

// WithDecompile adds a decompile to the builder
func (app *assignableResourceBuilder) WithDecompile(decompile string) AssignableResourceBuilder {
	app.decompile = decompile
	return app
}

// WihAmountByQuery adds an amountByQuery to the builder
func (app *assignableResourceBuilder) WihAmountByQuery(amountByQuery BytesReference) AssignableResourceBuilder {
	app.amountByQuery = amountByQuery
	return app
}

// WithRetrieveByQuery adds a retrieveByQuery to the builder
func (app *assignableResourceBuilder) WithRetrieveByQuery(retrieveByQuery BytesReference) AssignableResourceBuilder {
	app.retrieveByQuery = retrieveByQuery
	return app
}

// WithRetrieveByHash adds a retrieveByHash to the builder
func (app *assignableResourceBuilder) WithRetrieveByHash(retrieveByHash BytesReference) AssignableResourceBuilder {
	app.retrieveByHash = retrieveByHash
	return app
}

// IsAmount flags the builder as an amount
func (app *assignableResourceBuilder) IsAmount() AssignableResourceBuilder {
	app.isAmount = true
	return app
}

// Now builds a new AssignableResource instance
func (app *assignableResourceBuilder) Now() (AssignableResource, error) {
	data := [][]byte{}
	if app.compile != nil {
		data = append(data, []byte("compile"))
		data = append(data, app.compile.Hash().Bytes())
	}

	if app.decompile != "" {
		data = append(data, []byte("decompile"))
		data = append(data, []byte(app.decompile))
	}

	if app.amountByQuery != nil {
		data = append(data, []byte("amountByQuery"))
		data = append(data, app.amountByQuery.Hash().Bytes())
	}

	if app.retrieveByQuery != nil {
		data = append(data, []byte("retrieveByQuery"))
		data = append(data, app.retrieveByQuery.Hash().Bytes())
	}

	if app.retrieveByHash != nil {
		data = append(data, []byte("retrieveByHash"))
		data = append(data, app.retrieveByHash.Hash().Bytes())
	}

	if app.isAmount {
		data = append(data, []byte("isAmount"))
	}

	if len(data) <= 0 {
		return nil, errors.New("the AssignableResource is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.compile != nil {
		return createAssignableResourceWithCompile(*pHash, app.compile), nil
	}

	if app.decompile != "" {
		return createAssignableResourceWithDecompile(*pHash, app.decompile), nil
	}

	if app.amountByQuery != nil {
		return createAssignableResourceWithAmountByQuery(*pHash, app.amountByQuery), nil
	}

	if app.retrieveByQuery != nil {
		return createAssignableResourceWithRetrieveByQuery(*pHash, app.retrieveByQuery), nil
	}

	if app.retrieveByHash != nil {
		return createAssignableResourceWithRetrieveByHash(*pHash, app.retrieveByHash), nil
	}

	return createAssignableResourceWithAmount(*pHash), nil
}
