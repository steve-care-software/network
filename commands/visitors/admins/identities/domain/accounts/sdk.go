package accounts

import (
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/encryptors"
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/profiles"
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/signers"
	"steve.care/network/commands/visitors/stencils/domain/layers"
)

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithRoot(root layers.Layer) Builder
	WithProfile(profile profiles.Profile) Builder
	WithEncryptor(encryptor encryptors.Encryptor) Builder
	WithSigner(signer signers.Signer) Builder
	Now() (Account, error)
}

// Account represents the identity account
type Account interface {
	Root() layers.Layer
	Profile() profiles.Profile
	Encryptor() encryptors.Encryptor
	Signer() signers.Signer
}
