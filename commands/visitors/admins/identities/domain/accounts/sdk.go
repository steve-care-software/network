package accounts

import (
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/encryptors"
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/profiles"
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/signers"
)

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithProfile(profile profiles.Profile) Builder
	WithEncryptor(encryptor encryptors.Encryptor) Builder
	WithSigner(signer signers.Signer) Builder
	Now() (Account, error)
}

// Account represents the identity account
type Account interface {
	Profile() profiles.Profile
	Encryptor() encryptors.Encryptor
	Signer() signers.Signer
}
