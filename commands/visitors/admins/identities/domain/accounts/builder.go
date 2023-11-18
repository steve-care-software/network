package accounts

import (
	"errors"

	"steve.care/network/commands/visitors/admins/identities/domain/accounts/encryptors"
	"steve.care/network/commands/visitors/admins/identities/domain/accounts/signers"
)

type builder struct {
	root      []string
	encryptor encryptors.Encryptor
	signer    signers.Signer
}

func createBuilder() Builder {
	out := builder{
		root:      nil,
		encryptor: nil,
		signer:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root []string) Builder {
	app.root = root
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

// Now builds a new Account instance
func (app *builder) Now() (Account, error) {
	if app.root != nil && len(app.root) <= 0 {
		app.root = nil
	}

	if app.root == nil {
		return nil, errors.New("the root layer path is mandatory in order to build an Account instance")
	}

	if app.encryptor == nil {
		return nil, errors.New("the encryptor is mandatory in order to build an Account instance")
	}

	if app.signer == nil {
		return nil, errors.New("the signer is mandatory in order to build an Account instance")
	}

	return createAccount(
		app.root,
		app.encryptor,
		app.signer,
	), nil
}
