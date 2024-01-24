package groups

import (
	"errors"

	"steve.care/network/domain/schemas/groups/resources"
)

type elementBuilder struct {
	groups    Groups
	resources resources.Resources
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		groups:    nil,
		resources: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithGroups add groups to the builder
func (app *elementBuilder) WithGroups(groups Groups) ElementBuilder {
	app.groups = groups
	return app
}

// WithResources add resources to the builder
func (app *elementBuilder) WithResources(resources resources.Resources) ElementBuilder {
	app.resources = resources
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.groups != nil {
		return createElementWithGroups(app.groups), nil
	}

	if app.resources != nil {
		return createElementWithResources(app.resources), nil
	}

	return nil, errors.New("the Element is invalid")
}
