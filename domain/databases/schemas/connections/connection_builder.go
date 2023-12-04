package connections

import (
	"errors"

	"steve.care/network/domain/databases/schemas/connections/pointers"
)

type connectionBuilder struct {
	from pointers.Pointer
	to   pointers.Pointer
}

func createConnectionBuilder() ConnectionBuilder {
	out := connectionBuilder{
		from: nil,
		to:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionBuilder) Create() ConnectionBuilder {
	return createConnectionBuilder()
}

// From adds a from pointer to the builder
func (app *connectionBuilder) From(from pointers.Pointer) ConnectionBuilder {
	app.from = from
	return app
}

// To adds a to pointer to the builder
func (app *connectionBuilder) To(to pointers.Pointer) ConnectionBuilder {
	app.to = to
	return app
}

// Now builds a new Connection instance
func (app *connectionBuilder) Now() (Connection, error) {
	if app.from == nil {
		return nil, errors.New("the from pointer is mandatory in order to build a Connection")
	}

	if app.to == nil {
		return nil, errors.New("the to pointer is mandatory in order to build a Connection")
	}

	return createConnection(
		app.from,
		app.to,
	), nil
}
