package accounts

import (
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/accounts/workspaces"
)

type account struct {
	username  string
	encryptor encryptors.Encryptor
	signer    signers.Signer
	workspace workspaces.Workspace
}

func createAccount(
	username string,
	encryptor encryptors.Encryptor,
	signer signers.Signer,
) Account {
	return createAccountInternally(username, encryptor, signer, nil)
}

func createAccountWithWorkspace(
	username string,
	encryptor encryptors.Encryptor,
	signer signers.Signer,
	workspace workspaces.Workspace,
) Account {
	return createAccountInternally(username, encryptor, signer, workspace)
}

func createAccountInternally(
	username string,
	encryptor encryptors.Encryptor,
	signer signers.Signer,
	workspace workspaces.Workspace,
) Account {
	out := account{
		username:  username,
		encryptor: encryptor,
		signer:    signer,
		workspace: workspace,
	}

	return &out
}

// Username returns the username
func (obj *account) Username() string {
	return obj.username
}

// Encryptor returns the encryptor
func (obj *account) Encryptor() encryptors.Encryptor {
	return obj.encryptor
}

// Signer returns the signer
func (obj *account) Signer() signers.Signer {
	return obj.signer
}

// HasWorkspace returns true if there is a workspace, false otherwise
func (obj *account) HasWorkspace() bool {
	return obj.workspace != nil
}

// Workspace returns the workspace, if any
func (obj *account) Workspace() workspaces.Workspace {
	return obj.workspace
}
