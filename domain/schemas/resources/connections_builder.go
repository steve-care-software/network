package resources

import "errors"

type connectionsBuilder struct {
	list []Connection
}

func createConnectionsBuilder() ConnectionsBuilder {
	out := connectionsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionsBuilder) Create() ConnectionsBuilder {
	return createConnectionsBuilder()
}

// WithList adds a list to the builder
func (app *connectionsBuilder) WithList(list []Connection) ConnectionsBuilder {
	app.list = list
	return app
}

// Now builds a new Connections instance
func (app *connectionsBuilder) Now() (Connections, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Connection in order to build a Connections instance")
	}

	return createConnections(app.list), nil
}
