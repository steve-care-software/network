package applications

import (
	"errors"

	"steve.care/network/applications/applications/accounts"
	"steve.care/network/applications/applications/authenticates"
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/databases/transactions"
)

type builder struct {
	accAppBuilder  accounts.Builder
	authAppBuilder authenticates.Builder
	query          queries.Query
	trx            transactions.Transaction
	bitrate        int
}

func createBuilder(
	accAppBuilder accounts.Builder,
	authAppBuilder authenticates.Builder,
) Builder {
	out := builder{
		accAppBuilder:  accAppBuilder,
		authAppBuilder: authAppBuilder,
		query:          nil,
		trx:            nil,
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

// WithQuery adds a query to the builder
func (app *builder) WithQuery(query queries.Query) Builder {
	app.query = query
	return app
}

// WithTransaction adds a trx to the builder
func (app *builder) WithTransaction(trx transactions.Transaction) Builder {
	app.trx = trx
	return app
}

// WithBitrate adds a bitrate to the builder
func (app *builder) WithBitrate(bitrate int) Builder {
	app.bitrate = bitrate
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.query == nil {
		return nil, errors.New("the query is mandatory in order to build an Application instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transaction is mandatory in order to build an Application instance")
	}

	if app.bitrate <= 0 {
		return nil, errors.New("the bitrate is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.accAppBuilder,
		app.authAppBuilder,
		app.trx,
		app.query,
		app.bitrate,
	), nil
}
