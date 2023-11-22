package layers

import (
	"steve.care/network/libraries/hash"
)

type bytesReferencesBuilder struct {
	hashAdapter hash.Adapter
	list        []BytesReference
}

func createBytesReferencesBuilder(
	hashAdapter hash.Adapter,
) BytesReferencesBuilder {
	out := bytesReferencesBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *bytesReferencesBuilder) Create() BytesReferencesBuilder {
	return createBytesReferencesBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *bytesReferencesBuilder) WithList(list []BytesReference) BytesReferencesBuilder {
	app.list = list
	return app
}

// Now builds a new BytesReferences instance
func (app *bytesReferencesBuilder) Now() (BytesReferences, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	data := [][]byte{}
	for _, oneIns := range app.list {
		data = append(data, oneIns.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createBytesReferences(*pHash, app.list), nil
}
