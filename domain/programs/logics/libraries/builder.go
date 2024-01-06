package libraries

import (
	"errors"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
)

type builder struct {
	hashAdapter hash.Adapter
	layers      layers.Layers
	links       links.Links
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		layers:      nil,
		links:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithLayers add layers to the builder
func (app *builder) WithLayers(layers layers.Layers) Builder {
	app.layers = layers
	return app
}

// WithLinks add links to the builder
func (app *builder) WithLinks(links links.Links) Builder {
	app.links = links
	return app
}

// Now builds a new Library instance
func (app *builder) Now() (Library, error) {
	data := [][]byte{}
	if app.layers != nil {
		data = append(data, app.layers.Hash().Bytes())
	}

	if app.links != nil {
		data = append(data, app.links.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Library is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.layers != nil && app.links != nil {
		return createLibraryWithLayersAndLinks(*pHash, app.layers, app.links), nil
	}

	if app.layers != nil {
		return createLibraryWithLayers(*pHash, app.layers), nil
	}

	return createLibraryWithLinks(*pHash, app.links), nil
}
