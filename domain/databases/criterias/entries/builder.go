package entries

import (
	"errors"

	"steve.care/network/domain/databases/criterias/entries/resources"
)

type builder struct {
	resource resources.Resource
	fields   []string
}

func createBuilder() Builder {
	out := builder{
		resource: nil,
		fields:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithResource adds a resource to the builder
func (app *builder) WithResource(resource resources.Resource) Builder {
	app.resource = resource
	return app
}

// WithFields adds a fields to the builder
func (app *builder) WithFields(fields []string) Builder {
	app.fields = fields
	return app
}

// Now builds a new Entry instance
func (app *builder) Now() (Entry, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build an Entry instance")
	}

	if app.fields != nil && len(app.fields) <= 0 {
		app.fields = nil
	}

	if app.fields == nil {
		return nil, errors.New("the fields is mandatory in order to build an Entry instance")
	}

	return createEntry(
		app.resource,
		app.fields,
	), nil
}
