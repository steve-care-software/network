package resources

import "errors"

type connectionBuilder struct {
	field     string
	reference Pointer
}

func createConnectionBuilder() ConnectionBuilder {
	out := connectionBuilder{
		field:     "",
		reference: nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionBuilder) Create() ConnectionBuilder {
	return createConnectionBuilder()
}

// WithField adds a field to the builder
func (app *connectionBuilder) WithField(field string) ConnectionBuilder {
	app.field = field
	return app
}

// WithReference adds a reference to the builder
func (app *connectionBuilder) WithReference(reference Pointer) ConnectionBuilder {
	app.reference = reference
	return app
}

// Now builds a new Connection instance
func (app *connectionBuilder) Now() (Connection, error) {
	if app.field == "" {
		return nil, errors.New("the field is mandatory in order to build a Connection instance")
	}

	if app.reference == nil {
		return nil, errors.New("the reference is mandatory in order to build a Connection instance")
	}

	return createConnection(app.field, app.reference), nil
}
