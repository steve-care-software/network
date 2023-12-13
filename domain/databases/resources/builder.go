package resources

import (
	"errors"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/databases/resources/tokens"
	"steve.care/network/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	token       tokens.Token
	signature   signers.Signature
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		token:       nil,
		signature:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithToken adds a token to the builder
func (app *builder) WithToken(token tokens.Token) Builder {
	app.token = token
	return app
}

// WithSignature adds a signature to the builder
func (app *builder) WithSignature(signature signers.Signature) Builder {
	app.signature = signature
	return app
}

// Now builds a new Resource instance
func (app *builder) Now() (Resource, error) {
	if app.token == nil {
		return nil, errors.New("the token is mandatory in order to build a Resource instance")
	}

	if app.signature == nil {
		return nil, errors.New("the signature is mandatory in order to build a Resource instance")
	}

	sigBytes, err := app.signature.Bytes()
	if err != nil {
		return nil, err
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.token.Hash().Bytes(),
		sigBytes,
	})

	if err != nil {
		return nil, err
	}

	return createResource(*pHash, app.token, app.signature), nil
}
