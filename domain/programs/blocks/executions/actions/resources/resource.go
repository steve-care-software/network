package resources

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens"
)

type resource struct {
	hash      hash.Hash
	token     tokens.Token
	signature signers.Signature
}

func createResource(
	hash hash.Hash,
	token tokens.Token,
	signature signers.Signature,
) Resource {
	out := resource{
		hash:      hash,
		token:     token,
		signature: signature,
	}

	return &out
}

// Hash returns the hash
func (obj *resource) Hash() hash.Hash {
	return obj.hash
}

// Token returns the token
func (obj *resource) Token() tokens.Token {
	return obj.token
}

// Signature returns the signature
func (obj *resource) Signature() signers.Signature {
	return obj.signature
}
