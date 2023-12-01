package applications

import (
	"errors"

	"steve.care/network/applications/applications/accounts"
	"steve.care/network/applications/applications/authenticates"
	"steve.care/network/domain/databases"
)

type builder struct {
	accAppBuilder  accounts.Builder
	authAppBuilder authenticates.Builder
	database       databases.Database
	bitrate        int
}

func createBuilder(
	accAppBuilder accounts.Builder,
	authAppBuilder authenticates.Builder,
) Builder {
	out := builder{
		accAppBuilder:  accAppBuilder,
		authAppBuilder: authAppBuilder,
		database:       nil,
		bitrate:        0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.accAppBuilder,
		app.authAppBuilder,
	)
}

// WithDatabase adds a database to the builder
func (app *builder) WithDatabase(database databases.Database) Builder {
	app.database = database
	return app
}

// WithBitrate adds a bitrate to the builder
func (app *builder) WithBitrate(bitrate int) Builder {
	app.bitrate = bitrate
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.database == nil {
		return nil, errors.New("the database is mandatory in order to build an Application instance")
	}

	if app.bitrate <= 0 {
		return nil, errors.New("the bitrate is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.accAppBuilder,
		app.authAppBuilder,
		app.database,
		app.bitrate,
	), nil
}
