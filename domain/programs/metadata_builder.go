package programs

import (
	"errors"

	"steve.care/network/domain/hash"
)

type metadataBuilder struct {
	hashAdapter hash.Adapter
	name        string
	parent      Program
}

func createMetaDataBuilder(
	hashAdapter hash.Adapter,
) MetaDataBuilder {
	out := metadataBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		parent:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *metadataBuilder) Create() MetaDataBuilder {
	return createMetaDataBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *metadataBuilder) WithName(name string) MetaDataBuilder {
	app.name = name
	return app
}

// WithParent adds a parent to the builder
func (app *metadataBuilder) WithParent(parent Program) MetaDataBuilder {
	app.parent = parent
	return app
}

// Now builds a new MetaData instance
func (app *metadataBuilder) Now() (MetaData, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a MetaData instance")
	}

	if app.parent == nil {
		return nil, errors.New("the parent is mandatory in order to build a MetaData instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.parent.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createMetaData(*pHash, app.name, app.parent), nil
}
