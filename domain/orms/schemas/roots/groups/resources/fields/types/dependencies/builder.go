package dependencies

import "errors"

type builder struct {
	retriever string
	groups    []string
	resource  string
	pKind     *uint8
}

func createBuilder() Builder {
	out := builder{
		retriever: "",
		groups:    nil,
		resource:  "",
		pKind:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRetriever adds a retriever to the builder
func (app *builder) WithRetriever(retriever string) Builder {
	app.retriever = retriever
	return app
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

// WithKind adds a kind to the builder
func (app *builder) WithKind(kind uint8) Builder {
	app.pKind = &kind
	return app
}

// Now builds a new Dependency instance
func (app *builder) Now() (Dependency, error) {
	if app.retriever == "" {
		return nil, errors.New("the retriever is mandatory in order to build a Dependency instance")
	}

	if app.groups != nil && len(app.groups) <= 0 {
		app.groups = nil
	}

	if app.groups == nil {
		return nil, errors.New("the groups is mandatory in order to build a Dependency instance")
	}

	if app.resource == "" {
		return nil, errors.New("the resource is mandatory in order to build a Dependency instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Dependency instance")
	}

	return createDependency(app.retriever, app.groups, app.resource, *app.pKind), nil
}
