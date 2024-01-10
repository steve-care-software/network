package receipts

import (
	"errors"

	"steve.care/network/domain/receipts"
	"steve.care/network/domain/receipts/commands"
	"steve.care/network/domain/receipts/commands/results"
)

type builder struct {
	receipt receipts.Receipt
	command commands.Command
	result  results.Result
	success results.Success
	failure results.Failure
	link    commands.Link
}

func createBuilder() Builder {
	out := builder{
		receipt: nil,
		command: nil,
		result:  nil,
		success: nil,
		failure: nil,
		link:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithReceipt adds a receipt to the builder
func (app *builder) WithReceipt(receipt receipts.Receipt) Builder {
	app.receipt = receipt
	return app
}

// WithCommand adds a command to the builder
func (app *builder) WithCommand(command commands.Command) Builder {
	app.command = command
	return app
}

// WithResult adds a result to the builder
func (app *builder) WithResult(result results.Result) Builder {
	app.result = result
	return app
}

// WithSuccess adds a success to the builder
func (app *builder) WithSuccess(success results.Success) Builder {
	app.success = success
	return app
}

// WithFailure adds a failure to the builder
func (app *builder) WithFailure(failure results.Failure) Builder {
	app.failure = failure
	return app
}

// WithLink adds a link to the builder
func (app *builder) WithLink(link commands.Link) Builder {
	app.link = link
	return app
}

// Now builds a new Receipt instance
func (app *builder) Now() (Receipt, error) {
	if app.receipt != nil {
		return createReceiptWithReceipt(app.receipt), nil
	}

	if app.command != nil {
		return createReceiptWithCommand(app.command), nil
	}

	if app.result != nil {
		return createReceiptWithResult(app.result), nil
	}

	if app.success != nil {
		return createReceiptWithSuccess(app.success), nil
	}

	if app.failure != nil {
		return createReceiptWithFailure(app.failure), nil
	}

	if app.link != nil {
		return createReceiptWithLink(app.link), nil
	}

	return nil, errors.New("the Receipt resource is invalid")
}
