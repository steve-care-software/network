package layers

import (
	"errors"

	"steve.care/network/domain/hash"
)

type assignableResourceBuilder struct {
	hashAdapter     hash.Adapter
	compile         string
	decompile       string
	amountByQuery   string
	listByQuery     string
	retrieveByQuery string
	retrieveByHash  string
	isAmount        bool
}

func createAssignableResourceBuilder(
	hashAdapter hash.Adapter,
) AssignableResourceBuilder {
	out := assignableResourceBuilder{
		hashAdapter:     hashAdapter,
		compile:         "",
		decompile:       "",
		amountByQuery:   "",
		listByQuery:     "",
		retrieveByQuery: "",
		retrieveByHash:  "",
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
func (app *assignableResourceBuilder) WithCompile(compile string) AssignableResourceBuilder {
	app.compile = compile
	return app
}

// WithDecompile adds a decompile to the builder
func (app *assignableResourceBuilder) WithDecompile(decompile string) AssignableResourceBuilder {
	app.decompile = decompile
	return app
}

// WihAmountByQuery adds an amountByQuery to the builder
func (app *assignableResourceBuilder) WihAmountByQuery(amountByQuery string) AssignableResourceBuilder {
	app.amountByQuery = amountByQuery
	return app
}

// WithListByQuery adds a listByQuery to the builder
func (app *assignableResourceBuilder) WithListByQuery(listByQuery string) AssignableResourceBuilder {
	app.listByQuery = listByQuery
	return app
}

// WithRetrieveByQuery adds a retrieveByQuery to the builder
func (app *assignableResourceBuilder) WithRetrieveByQuery(retrieveByQuery string) AssignableResourceBuilder {
	app.retrieveByQuery = retrieveByQuery
	return app
}

// WithRetrieveByHash adds a retrieveByHash to the builder
func (app *assignableResourceBuilder) WithRetrieveByHash(retrieveByHash string) AssignableResourceBuilder {
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
	if app.compile != "" {
		data = append(data, []byte("compile"))
		data = append(data, []byte(app.compile))
	}

	if app.decompile != "" {
		data = append(data, []byte("decompile"))
		data = append(data, []byte(app.decompile))
	}

	if app.amountByQuery != "" {
		data = append(data, []byte("amountByQuery"))
		data = append(data, []byte(app.amountByQuery))
	}

	if app.listByQuery != "" {
		data = append(data, []byte("listByQuery"))
		data = append(data, []byte(app.listByQuery))
	}

	if app.retrieveByQuery != "" {
		data = append(data, []byte("retrieveByQuery"))
		data = append(data, []byte(app.retrieveByQuery))
	}

	if app.retrieveByHash != "" {
		data = append(data, []byte("retrieveByHash"))
		data = append(data, []byte(app.retrieveByHash))
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

	if app.compile != "" {
		return createAssignableResourceWithCompile(*pHash, app.compile), nil
	}

	if app.decompile != "" {
		return createAssignableResourceWithDecompile(*pHash, app.decompile), nil
	}

	if app.amountByQuery != "" {
		return createAssignableResourceWithAmountByQuery(*pHash, app.amountByQuery), nil
	}

	if app.listByQuery != "" {
		return createAssignableResourceWithListByQuery(*pHash, app.listByQuery), nil
	}

	if app.retrieveByQuery != "" {
		return createAssignableResourceWithRetrieveByQuery(*pHash, app.retrieveByQuery), nil
	}

	if app.retrieveByHash != "" {
		return createAssignableResourceWithRetrieveByHash(*pHash, app.retrieveByHash), nil
	}

	return createAssignableResourceWithAmount(*pHash), nil
}
