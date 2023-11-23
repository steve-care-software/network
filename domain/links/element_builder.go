package links

import (
	"errors"
	"strings"

	"steve.care/network/domain/hash"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	container   []string
	condition   Condition
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		container:   nil,
		condition:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(
		app.hashAdapter,
	)
}

// WithContainer adds a container to the builder
func (app *elementBuilder) WithContainer(container []string) ElementBuilder {
	app.container = container
	return app
}

// WithCondition adds a condition to the builder
func (app *elementBuilder) WithCondition(condition Condition) ElementBuilder {
	app.condition = condition
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.container != nil && len(app.container) <= 0 {
		app.container = nil
	}

	if app.container == nil {
		return nil, errors.New("the container is mandatory in order to build an Element instance")
	}

	data := [][]byte{
		[]byte(strings.Join(app.container, "/")),
	}

	if app.condition != nil {
		data = append(data, app.condition.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.condition != nil {
		return createElementWithCondition(*pHash, app.container, app.condition), nil
	}

	return createElement(*pHash, app.container), nil
}
