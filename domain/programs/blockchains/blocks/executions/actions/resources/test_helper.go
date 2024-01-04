package resources

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources/tokens"
)

// NewResourceForTests creates a new resource for tests
func NewResourceForTests(input tokens.Token) Resource {
	msg := input.Hash().Bytes()
	signature, err := signers.NewFactory().Create().Sign(msg)
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithToken(input).WithSignature(signature).Now()
	if err != nil {
		panic(err)
	}

	return ins

}
