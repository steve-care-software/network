package accounts

import (
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
)

type account struct {
	username  string
	root      []string
	encryptor encryptors.Encryptor
	signer    signers.Signer
}

func createAccount(
	username string,
	root []string,
	encryptor encryptors.Encryptor,
	signer signers.Signer,
) Account {
	out := account{
		username:  username,
		root:      root,
		encryptor: encryptor,
		signer:    signer,
	}

	return &out
}

// Username returns the username
func (obj *account) Username() string {
	return obj.username
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
