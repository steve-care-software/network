package receipts

import (
	"errors"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands"
)

type builder struct {
	hashAdapter hash.Adapter
	commands    commands.Commands
	signature   signers.Signature
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		commands:    nil,
		signature:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithCommands add commands to the builder
func (app *builder) WithCommands(commands commands.Commands) Builder {
	app.commands = commands
	return app
}

// WithSignature add signature to the builder
func (app *builder) WithSignature(signature signers.Signature) Builder {
	app.signature = signature
	return app
}

// Now builds a new Receipt instance
func (app *builder) Now() (Receipt, error) {
	if app.commands == nil {
		return nil, errors.New("the commands is mandatory in order to build a Receipt instance")
	}

	if app.signature == nil {
		return nil, errors.New("the signature is mandatory in order to build a Receipt instance")
	}

	sigBytes, err := app.signature.Bytes()
	if err != nil {
		return nil, err
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.commands.Hash().Bytes(),
		sigBytes,
	})

	if err != nil {
		return nil, err
	}

	return createReceipt(*pHash, app.commands, app.signature), nil
}
