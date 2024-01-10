package tokens

import (
	"time"

	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/dashboards"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/links"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/queries"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/receipts"
)

// NewTokenWithLayerForTests creates a new token with layer for tests
func NewTokenWithLayerForTests(input layers.Layer) Token {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithLayer(input).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTokenWithLinkForTests creates a new token with link for tests
func NewTokenWithLinkForTests(input links.Link) Token {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithLink(input).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTokenWithReceiptForTests creates a new token with receipt for tests
func NewTokenWithReceiptForTests(input receipts.Receipt) Token {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithReceipt(input).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTokenWithQueryForTests creates a new token with query for tests
func NewTokenWithQueryForTests(input queries.Query) Token {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithQuery(input).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTokenWithDashboardForTests creates a new token with dashboard for tests
func NewTokenWithDashboardForTests(input dashboards.Dashboard) Token {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithDashboard(input).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
