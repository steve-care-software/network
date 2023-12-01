package jsons

import (
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
)

// Account represents an account
type Account struct {
	Username  string `json:"username"`
	Encryptor []byte `json:"encryptor"`
	Signer    []byte `json:"signer"`
}

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
