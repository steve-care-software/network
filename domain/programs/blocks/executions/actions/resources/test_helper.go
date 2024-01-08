package resources

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens"
)

// NewResourceForTests creates a new resource for tests
func NewResourceForTests(token tokens.Token, signature signers.Signature) Resource {
	ins, err := NewBuilder().Create().WithToken(token).WithSignature(signature).Now()
	if err != nil {
		panic(err)
	}

	return ins

}
