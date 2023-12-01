package authenticates

import (
	"errors"

	"steve.care/network/applications/applications/authenticates/commands"
	"steve.care/network/applications/applications/authenticates/layers"
	"steve.care/network/applications/applications/authenticates/links"
	"steve.care/network/applications/applications/authenticates/receipts"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases"
)

type builder struct {
	commandBuilder commands.Builder
	layerBuilder   layers.Builder
	linkBuilder    links.Builder
	receiptBuilder receipts.Builder
	trx            databases.Transaction
	query          databases.Query
	credentials    credentials.Credentials
}

func createBuilder(
	commandBuilder commands.Builder,
	layerBuilder layers.Builder,
	linkBuilder links.Builder,
	receiptBuilder receipts.Builder,
) Builder {
	out := builder{
		commandBuilder: commandBuilder,
		layerBuilder:   layerBuilder,
		linkBuilder:    linkBuilder,
		receiptBuilder: receiptBuilder,
		trx:            nil,
		query:          nil,
		credentials:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.commandBuilder,
		app.layerBuilder,
		app.linkBuilder,
		app.receiptBuilder,
	)
}

// WithCredentials add credentials to the builder
func (app *builder) WithCredentials(credentials credentials.Credentials) Builder {
	app.credentials = credentials
	return app
}

// WithQuery add a query to the builder
func (app *builder) WithQuery(query databases.Query) Builder {
	app.query = query
	return app
}

// WithTransaction adds a transaction to the builder
func (app *builder) WithTransaction(trx databases.Transaction) Builder {
	app.trx = trx
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.credentials == nil {
		return nil, errors.New("the credentials is mandatory in order to build an Application instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transaction is mandatory in order to build an Application instance")
	}

	if app.query == nil {
		return nil, errors.New("the query is mandatory in order to build an Application instance")
	}

	commandsApp, err := app.commandBuilder.Create().
		WithCredentials(app.credentials).
		Now()

	if err != nil {
		return nil, err
	}

	layersApp, err := app.layerBuilder.Create().
		WithCredentials(app.credentials).
		Now()

	if err != nil {
		return nil, err
	}

	linksApp, err := app.linkBuilder.Create().
		WithCredentials(app.credentials).
		Now()

	if err != nil {
		return nil, err
	}

	receiptApp, err := app.receiptBuilder.Create().
		WithCredentials(app.credentials).
		Now()

	if err != nil {
		return nil, err
	}

	return createApplication(
		receiptApp,
		layersApp,
		linksApp,
		commandsApp,
	), nil
}