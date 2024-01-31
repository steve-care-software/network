package sqllites

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
)

// NewResourceFromTokenForTests creates a new resource from token for tests
func NewResourceFromTokenForTests(token tokens.Token) resources.Resource {
	msg := token.Hash().Bytes()
	signature, err := signers.NewFactory().Create().Sign(msg)
	if err != nil {
		panic(err)
	}

	return resources.NewResourceForTests(token, signature)
}
