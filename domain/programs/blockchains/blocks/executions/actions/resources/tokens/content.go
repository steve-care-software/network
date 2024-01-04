package tokens

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources/tokens/links"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources/tokens/queries"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources/tokens/receipts"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources/tokens/suites"
)

type content struct {
	layer   layers.Layer
	link    links.Link
	suite   suites.Suite
	receipt receipts.Receipt
	query   queries.Query
}

func createContentWithLayer(
	layer layers.Layer,
) Content {
	return createContentInternally(
		layer,
		nil,
		nil,
		nil,
		nil,
	)
}

func createContentWithLink(
	link links.Link,
) Content {
	return createContentInternally(
		nil,
		link,
		nil,
		nil,
		nil,
	)
}

func createContentWithSuite(
	suite suites.Suite,
) Content {
	return createContentInternally(
		nil,
		nil,
		suite,
		nil,
		nil,
	)
}

func createContentWithReceipt(
	receipt receipts.Receipt,
) Content {
	return createContentInternally(
		nil,
		nil,
		nil,
		receipt,
		nil,
	)
}

func createContentWithQuery(
	query queries.Query,
) Content {
	return createContentInternally(
		nil,
		nil,
		nil,
		nil,
		query,
	)
}

func createContentInternally(
	layer layers.Layer,
	link links.Link,
	suite suites.Suite,
	receipt receipts.Receipt,
	query queries.Query,
) Content {
	out := content{
		layer:   layer,
		link:    link,
		suite:   suite,
		receipt: receipt,
		query:   query,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return nil
}

// IsLayer returns true if there is a layer, false otherwise
func (obj *content) IsLayer() bool {
	return obj.layer != nil
}

// Layer returns the layer, if any
func (obj *content) Layer() layers.Layer {
	return obj.layer
}

// IsLink returns true if there is a link, false otherwise
func (obj *content) IsLink() bool {
	return obj.link != nil
}

// Link returns the link, if any
func (obj *content) Link() links.Link {
	return obj.link
}

// IsSuite returns true if there is a suite, false otherwise
func (obj *content) IsSuite() bool {
	return obj.suite != nil
}

// Suite returns the suite, if any
func (obj *content) Suite() suites.Suite {
	return obj.suite
}

// IsReceipt returns true if there is a receipt, false otherwise
func (obj *content) IsReceipt() bool {
	return obj.receipt != nil
}

// Receipt returns the receipt, if any
func (obj *content) Receipt() receipts.Receipt {
	return obj.receipt
}

// IsQuery returns true if there is a query, false otherwise
func (obj *content) IsQuery() bool {
	return obj.query != nil
}

// Query returns the query, if any
func (obj *content) Query() queries.Query {
	return obj.query
}
