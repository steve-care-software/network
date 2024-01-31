package groups

import (
	"errors"

	"steve.care/network/domain/schemas/roots/groups/resources"
)

type elementBuilder struct {
	group    Group
	resource resources.Resource
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		group:    nil,
		resource: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithGroup adds a group to the builder
func (app *elementBuilder) WithGroup(group Group) ElementBuilder {
	app.group = group
	return app
}

// WithResource adds a resource to the builder
func (app *elementBuilder) WithResource(resource resources.Resource) ElementBuilder {
	app.resource = resource
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.group != nil {
		return createElementWithGroup(app.group), nil
	}

	if app.resource != nil {
		return createElementWithResource(app.resource), nil
	}

	return nil, errors.New("the Element is invalid")
}
