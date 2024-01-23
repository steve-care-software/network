package executions

import (
	"errors"
	"strconv"
	"time"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/threads"
	"steve.care/network/domain/receipts"
)

type builder struct {
	hashAdapter hash.Adapter
	thread      threads.Thread
	receipt     receipts.Receipt
	pBeginsOn   *time.Time
	pEndsOn     *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		thread:      nil,
		receipt:     nil,
		pBeginsOn:   nil,
		pEndsOn:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithThread adds a thread to the builder
func (app *builder) WithThread(thread threads.Thread) Builder {
	app.thread = thread
	return app
}

// WithReceipt adds a receipt to the builder
func (app *builder) WithReceipt(receipt receipts.Receipt) Builder {
	app.receipt = receipt
	return app
}

// BeginsOn adds a beginsOn to the builder
func (app *builder) BeginsOn(beginsOn time.Time) Builder {
	app.pBeginsOn = &beginsOn
	return app
}

// EndsOn adds an endsOn to the builder
func (app *builder) EndsOn(endsOn time.Time) Builder {
	app.pEndsOn = &endsOn
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
	if app.thread == nil {
		return nil, errors.New("the thread is mandatory in order to build an Execution instance")
	}

	if app.receipt == nil {
		return nil, errors.New("the receipt is mandatory in order to build an Execution instance")
	}

	if app.pBeginsOn == nil {
		return nil, errors.New("the beginsOn is mandatory in order to build an Execution instance")
	}

	if app.pEndsOn == nil {
		return nil, errors.New("the endsOn is mandatory in order to build an Execution instance")
	}

	if app.pEndsOn.Before(*app.pBeginsOn) {
		return nil, errors.New("the endsOn time cannot be before the beginsOn time")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.thread.Hash().Bytes(),
		app.receipt.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.pBeginsOn.UnixNano()))),
		[]byte(strconv.Itoa(int(app.pEndsOn.UnixNano()))),
	})

	if err != nil {
		return nil, err
	}

	return createExecution(
		*pHash,
		app.thread,
		app.receipt,
		*app.pBeginsOn,
		*app.pEndsOn,
	), nil
}
