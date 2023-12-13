package tokens

import (
	"time"

	"steve.care/network/domain/databases/resources/tokens/layers"
	"steve.care/network/domain/databases/resources/tokens/links"
	"steve.care/network/domain/databases/resources/tokens/queries"
	"steve.care/network/domain/databases/resources/tokens/receipts"
	"steve.care/network/domain/databases/resources/tokens/suites"
	"steve.care/network/domain/hash"
)

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
}
