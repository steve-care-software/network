package tokens

import (
	"errors"
	"strconv"
	"time"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/links"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/queries"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/receipts"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/suites"
)

type builder struct {
	hashAdapter hash.Adapter
	layer       layers.Layer
	link        links.Link
	suite       suites.Suite
	receipt     receipts.Receipt
	query       queries.Query
	pCreatedOn  *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		layer:       nil,
		link:        nil,
		suite:       nil,
		receipt:     nil,
		query:       nil,
		pCreatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithLayer adds a layer to the builder
func (app *builder) WithLayer(layer layers.Layer) Builder {
	app.layer = layer
	return app
}

// WithLink adds a link to the builder
func (app *builder) WithLink(link links.Link) Builder {
	app.link = link
	return app
}

// WithSuite adds a suite to the builder
func (app *builder) WithSuite(suite suites.Suite) Builder {
	app.suite = suite
	return app
}

// WithReceipt adds a receipt to the builder
func (app *builder) WithReceipt(receipt receipts.Receipt) Builder {
	app.receipt = receipt
	return app
}

// WithQuery adds a query to the builder
func (app *builder) WithQuery(query queries.Query) Builder {
	app.query = query
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Token instance
func (app *builder) Now() (Token, error) {
	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Token instance")
	}

	data := [][]byte{
		[]byte(strconv.Itoa(app.pCreatedOn.UTC().Nanosecond())),
	}

	var content Content
	if app.layer != nil {
		content = createContentWithLayer(app.layer)
		data = append(data, app.layer.Hash().Bytes())
	}

	if app.link != nil {
		content = createContentWithLink(app.link)
		data = append(data, app.link.Hash().Bytes())
	}

	if app.suite != nil {
		content = createContentWithSuite(app.suite)
		data = append(data, app.suite.Hash().Bytes())
	}

	if app.receipt != nil {
		content = createContentWithReceipt(app.receipt)
		data = append(data, app.receipt.Hash().Bytes())
	}

	if app.query != nil {
		content = createContentWithQuery(app.query)
		data = append(data, app.query.Hash().Bytes())
	}

	if content == nil {
		return nil, errors.New("the Token is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createToken(
		*pHash,
		content,
		*app.pCreatedOn,
	), nil
}
