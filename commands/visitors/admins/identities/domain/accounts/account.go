package accounts

import (
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/encryptors"
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/signers"
)

type account struct {
	root      []string
	encryptor encryptors.Encryptor
	signer    signers.Signer
}

func createAccount(
	root []string,
	encryptor encryptors.Encryptor,
	signer signers.Signer,
) Account {
	out := account{
		root:      root,
		encryptor: encryptor,
		signer:    signer,
	}

	return &out
}

// Root returns the root
func (obj *account) Root() []string {
	return obj.root
}

// Encryptor returns the encryptor
func (obj *account) Encryptor() encryptors.Encryptor {
	return obj.encryptor
}

// Signer returns the signer
func (obj *account) Signer() signers.Signer {
	return obj.signer
}
