package accounts

import (
	"errors"

	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/accounts/workspaces"
)

type builder struct {
	username  string
	encryptor encryptors.Encryptor
	signer    signers.Signer
	workspace workspaces.Workspace
}

func createBuilder() Builder {
	out := builder{
		username:  "",
		encryptor: nil,
		signer:    nil,
		workspace: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithUsername adds a username to the builder
func (app *builder) WithUsername(username string) Builder {
	app.username = username
	return app
}

// WithEncryptor adds an encryptor to the builder
func (app *builder) WithEncryptor(encryptor encryptors.Encryptor) Builder {
	app.encryptor = encryptor
	return app
}

// WithSigner adds a signer to the builder
func (app *builder) WithSigner(signer signers.Signer) Builder {
	app.signer = signer
	return app
}

// WithWorkspace adds a workspace to the builder
func (app *builder) WithWorkspace(workspace workspaces.Workspace) Builder {
	app.workspace = workspace
	return app
}

// Now builds a new Account instance
func (app *builder) Now() (Account, error) {
	if app.username == "" {
		return nil, errors.New("the username is mandatory in order to build an Account instance")
	}

	if app.encryptor == nil {
		return nil, errors.New("the encryptor is mandatory in order to build an Account instance")
	}

	if app.signer == nil {
		return nil, errors.New("the signer is mandatory in order to build an Account instance")
	}

	if app.workspace != nil {
		return createAccountWithWorkspace(
			app.username,
			app.encryptor,
			app.signer,
			app.workspace,
		), nil
	}

	return createAccount(
		app.username,
		app.encryptor,
		app.signer,
	), nil
}
