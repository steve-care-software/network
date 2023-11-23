package accounts

import (
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithEncryptor(encryptor encryptors.Encryptor) Builder
	WithSigner(signer signers.Signer) Builder
	Now() (Account, error)
}

// Account represents the identity account
type Account interface {
	Username() string
	Encryptor() encryptors.Encryptor
	Signer() signers.Signer
}

// Repository represents the account repository
type Repository interface {
}

// Service represents the account service
type Service interface {
}
