package tokens

import (
	"time"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/dashboards"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/links"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/queries"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/receipts"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/suites"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	return createContentBuilder()
}

// Builder represents the token builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Token, error)
}

// Token represents the token
type Token interface {
	Hash() hash.Hash
	Content() Content
	CreatedOn() time.Time
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithLayer(layer layers.Layer) ContentBuilder
	WithLink(link links.Link) ContentBuilder
	WithSuite(suite suites.Suite) ContentBuilder
	WithReceipt(receipt receipts.Receipt) ContentBuilder
	WithQuery(query queries.Query) ContentBuilder
	WithDashboard(dashboard dashboards.Dashboard) ContentBuilder
	Now() (Content, error)
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
