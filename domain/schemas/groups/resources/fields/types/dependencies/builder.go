package dependencies

import "errors"

type builder struct {
	groups   []string
	resource string
}

func createBuilder() Builder {
	out := builder{
		groups:   nil,
		resource: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithGroups add groups to the builder
func (app *builder) WithGroups(groups []string) Builder {
	app.groups = groups
	return app
}

// WithResource add resource to the builder
func (app *builder) WithResource(resource string) Builder {
	app.resource = resource
	return app
}

// Now builds a new Dependency instance
func (app *builder) Now() (Dependency, error) {
	if app.groups != nil && len(app.groups) <= 0 {
		app.groups = nil
	}

	if app.groups == nil {
		return nil, errors.New("the groups is mandatory in order to build a Dependency instance")
	}

	if app.resource == "" {
		return nil, errors.New("the resource is mandatory in order to build a Dependency instance")
	}

	return createDependency(app.groups, app.resource), nil
}
