package conditions

import "errors"

type elementBuilder struct {
	condition Condition
	resource  Resource
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		condition: nil,
		resource:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithCondition adds a condition to the builder
func (app *elementBuilder) WithCondition(condition Condition) ElementBuilder {
	app.condition = condition
	return app
}

// WithResource adds a resource to the builder
func (app *elementBuilder) WithResource(resource Resource) ElementBuilder {
	app.resource = resource
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.condition != nil {
		return createElementWithCondition(app.condition), nil
	}

	if app.resource != nil {
		return createElementWithResource(app.resource), nil
	}

	return nil, errors.New("the Element is invalid")
}
