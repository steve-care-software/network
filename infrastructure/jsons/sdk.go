package jsons

import (
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
)

// NewAccountAdapter creates a new account adapter
func NewAccountAdapter() accounts.Adapter {
	builder := accounts.NewBuilder()
	encryptorAdapter := encryptors.NewAdapter()
	signerAdapter := signers.NewAdapter()
	return createAccountAdapter(
		builder,
		encryptorAdapter,
		signerAdapter,
	)
}
