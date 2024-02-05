package connections

import "errors"

type connectionBuilder struct {
	name string
	from []string
	to   []string
}

func createConnectionBuilder() ConnectionBuilder {
	out := connectionBuilder{
		name: "",
		from: nil,
		to:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionBuilder) Create() ConnectionBuilder {
	return createConnectionBuilder()
}

// WithName adds a name to the builder
func (app *connectionBuilder) WithName(name string) ConnectionBuilder {
	app.name = name
	return app
}

// From adds a from to the builder
func (app *connectionBuilder) From(from []string) ConnectionBuilder {
	app.from = from
	return app
}

// To adds a to to the builder
func (app *connectionBuilder) To(to []string) ConnectionBuilder {
	app.to = to
	return app
}

// Now builds a new Connection instance
func (app *connectionBuilder) Now() (Connection, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Connection instance")
	}

	if app.from != nil && len(app.from) <= 0 {
		app.from = nil
	}

	if app.from == nil {
		return nil, errors.New("the from reference is mandatory in order to build a Connection instance")
	}

	if app.to != nil && len(app.to) <= 0 {
		app.to = nil
	}

	if app.to == nil {
		return nil, errors.New("the to reference is mandatory in order to build a Connection instance")
	}

	return createConnection(
		app.name,
		app.from,
		app.to,
	), nil
}
