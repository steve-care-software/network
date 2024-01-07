package tokens

import (
	"time"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/dashboards"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/links"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/queries"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/receipts"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/suites"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the token builder
type Builder interface {
	Create() Builder
	WithLayer(layer layers.Layer) Builder
	WithLink(link links.Link) Builder
	WithSuite(suite suites.Suite) Builder
	WithReceipt(receipt receipts.Receipt) Builder
	WithQuery(query queries.Query) Builder
	WithDashboard(dashboard dashboards.Dashboard) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Token, error)
}

// Token represents the token
type Token interface {
	Hash() hash.Hash
	Content() Content
	CreatedOn() time.Time
}

// Content represents a resource content
type Content interface {
	Hash() hash.Hash
	IsLayer() bool
	Layer() layers.Layer
	IsLink() bool
	Link() links.Link
	IsSuite() bool
	Suite() suites.Suite
	IsReceipt() bool
	Receipt() receipts.Receipt
	IsQuery() bool
	Query() queries.Query
	IsDashboard() bool
	Dashboard() dashboards.Dashboard
}
