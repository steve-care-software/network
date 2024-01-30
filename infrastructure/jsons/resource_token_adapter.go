package jsons

import (
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	structs_tokens "steve.care/network/infrastructure/jsons/resources/tokens"
)

type resourceTokenAdapter struct {
	layerAdapter   *resourceTokenLayerAdapter
	linkAdapter    *resourceTokenLinkAdapter
	suiteAdapter   *resourceTokenSuiteAdapter
	receiptAdapter *resourceTokenReceiptAdapter
	queryAdapter   *resourceTokenQueryAdapter
	builder        tokens.Builder
	contentBuilder tokens.ContentBuilder
}

func (app *resourceTokenAdapter) toStruct(ins tokens.Token) (*structs_tokens.Token, error) {
	structContent := structs_tokens.Content{}
	content := ins.Content()
	if content.IsLayer() {
		layer := app.layerAdapter.toStruct(content.Layer())
		structContent.Layer = &layer
	}

	if content.IsLink() {
		link := app.linkAdapter.toStruct(content.Link())
		structContent.Link = &link
	}

	if content.IsSuite() {
		suite := app.suiteAdapter.toStruct(content.Suite())
		structContent.Suite = &suite
	}

	if content.IsReceipt() {
		receipt, err := app.receiptAdapter.toStruct(content.Receipt())
		if err != nil {
			return nil, err

		}

		structContent.Receipt = receipt
	}

	if content.IsQuery() {
		query := app.queryAdapter.toStruct(content.Query())
		structContent.Query = &query
	}

	return &structs_tokens.Token{
		Content:   structContent,
		CreatedOn: ins.CreatedOn(),
	}, nil
}

func (app *resourceTokenAdapter) toInstance(ins structs_tokens.Token) (tokens.Token, error) {
	content := ins.Content
	contentBuilder := app.contentBuilder.Create()
	if content.Layer != nil {
		layer, err := app.layerAdapter.toInstance(*content.Layer)
		if err != nil {
			return nil, err
		}

		contentBuilder.WithLayer(layer)
	}

	if content.Link != nil {
		link, err := app.linkAdapter.toInstance(*content.Link)
		if err != nil {
			return nil, err
		}

		contentBuilder.WithLink(link)
	}

	if content.Suite != nil {
		suite, err := app.suiteAdapter.toInstance(*content.Suite)
		if err != nil {
			return nil, err
		}

		contentBuilder.WithSuite(suite)
	}

	if content.Receipt != nil {
		receipt, err := app.receiptAdapter.toInstance(*content.Receipt)
		if err != nil {
			return nil, err
		}

		contentBuilder.WithReceipt(receipt)
	}

	if content.Query != nil {
		query, err := app.queryAdapter.toInstance(*content.Query)
		if err != nil {
			return nil, err
		}

		contentBuilder.WithQuery(query)
	}

	contentIns, err := contentBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		CreatedOn(ins.CreatedOn).
		WithContent(contentIns).
		Now()
}
