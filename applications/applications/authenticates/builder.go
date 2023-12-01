package authenticates

import (
	"errors"

	"steve.care/network/applications/applications/authenticates/commands"
	"steve.care/network/applications/applications/authenticates/layers"
	"steve.care/network/applications/applications/authenticates/links"
	"steve.care/network/applications/applications/authenticates/receipts"
	"steve.care/network/domain/credentials"
)

type builder struct {
	commandBuilder commands.Builder
	layerBuilder   layers.Builder
	linkBuilder    links.Builder
	receiptBuilder receipts.Builder
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

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.credentials == nil {
		return nil, errors.New("the credentials is mandatory in order to build an Application instance")
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
