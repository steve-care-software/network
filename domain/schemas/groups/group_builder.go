package groups

import "errors"

type groupBuilder struct {
	name     string
	elements Elements
}

func createGroupBuilder() GroupBuilder {
	out := groupBuilder{
		name:     "",
		elements: nil,
	}

	return &out
}

// Create initializes the builder
func (app *groupBuilder) Create() GroupBuilder {
	return createGroupBuilder()
}

// WithName adds a name to the builder
func (app *groupBuilder) WithName(name string) GroupBuilder {
	app.name = name
	return app
}

// WithElements add elements to the builder
func (app *groupBuilder) WithElements(elements Elements) GroupBuilder {
	app.elements = elements
	return app
}

// Now builds a new Group instance
func (app *groupBuilder) Now() (Group, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Group instance")
	}

	if app.elements == nil {
		return nil, errors.New("the elements is mandatory in order to build a Group instance")
	}

	return createGroup(app.name, app.elements), nil
}
