package tokens

import (
	"errors"

	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/dashboards"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/links"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/queries"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/receipts"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/suites"
)

type contentBuilder struct {
	layer     layers.Layer
	link      links.Link
	suite     suites.Suite
	receipt   receipts.Receipt
	query     queries.Query
	dashboard dashboards.Dashboard
}

func createContentBuilder() ContentBuilder {
	out := contentBuilder{
		layer:     nil,
		link:      nil,
		suite:     nil,
		receipt:   nil,
		query:     nil,
		dashboard: nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder()
}

// WithLayer adds a layer to the builder
func (app *contentBuilder) WithLayer(layer layers.Layer) ContentBuilder {
	app.layer = layer
	return app
}

// WithLink adds a link to the builder
func (app *contentBuilder) WithLink(link links.Link) ContentBuilder {
	app.link = link
	return app
}

// WithSuite adds a suite to the builder
func (app *contentBuilder) WithSuite(suite suites.Suite) ContentBuilder {
	app.suite = suite
	return app
}

// WithReceipt adds a receipt to the builder
func (app *contentBuilder) WithReceipt(receipt receipts.Receipt) ContentBuilder {
	app.receipt = receipt
	return app
}

// WithQuery adds a query to the builder
func (app *contentBuilder) WithQuery(query queries.Query) ContentBuilder {
	app.query = query
	return app
}

// WithDashboard adds a dashboard to the builder
func (app *contentBuilder) WithDashboard(dashboard dashboards.Dashboard) ContentBuilder {
	app.dashboard = dashboard
	return app
}

// Now builds a new Token instance
func (app *contentBuilder) Now() (Content, error) {
	if app.layer != nil {
		return createContentWithLayer(app.layer), nil
	}

	if app.link != nil {
		return createContentWithLink(app.link), nil
	}

	if app.suite != nil {
		return createContentWithSuite(app.suite), nil
	}

	if app.receipt != nil {
		return createContentWithReceipt(app.receipt), nil
	}

	if app.query != nil {
		return createContentWithQuery(app.query), nil
	}

	if app.dashboard != nil {
		return createContentWithDashboard(app.dashboard), nil
	}

	return nil, errors.New("the Token is invalid")
}
