package links

import (
	"errors"
	"strings"

	"steve.care/network/libraries/hash"
)

type originResourceBuilder struct {
	hashAdapter hash.Adapter
	container   []string
	isMandatory bool
}

func createOriginResourceBuilder(
	hashAdapter hash.Adapter,
) OriginResourceBuilder {
	out := originResourceBuilder{
		hashAdapter: hashAdapter,
		container:   nil,
		isMandatory: false,
	}

	return &out
}

// Create initializes the builder
func (app *originResourceBuilder) Create() OriginResourceBuilder {
	return createOriginResourceBuilder(
		app.hashAdapter,
	)
}

// WithContainer adds container to the builder
func (app *originResourceBuilder) WithContainer(container []string) OriginResourceBuilder {
	app.container = container
	return app
}

// IsMandatory flags the builder as mandatory
func (app *originResourceBuilder) IsMandatory() OriginResourceBuilder {
	app.isMandatory = true
	return app
}

// Now builds a new OriginResource instance
func (app *originResourceBuilder) Now() (OriginResource, error) {
	if app.container != nil && len(app.container) <= 0 {
		app.container = nil
	}

	if app.container == nil {
		return nil, errors.New("the container is mandatory in order to build an OriginResouce instance")
	}

	isMandatory := "false"
	if app.isMandatory {
		isMandatory = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strings.Join(app.container, "/")),
		[]byte(isMandatory),
	})

	if err != nil {
		return nil, err
	}

	return createOriginResource(*pHash, app.container, app.isMandatory), nil
}
