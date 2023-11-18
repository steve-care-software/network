package accounts

import (
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/encryptors"
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/signers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithRoot(root []string) Builder
	WithEncryptor(encryptor encryptors.Encryptor) Builder
	WithSigner(signer signers.Signer) Builder
	Now() (Account, error)
}

// Account represents the identity account
type Account interface {
	Root() []string
	Encryptor() encryptors.Encryptor
	Signer() signers.Signer
}
