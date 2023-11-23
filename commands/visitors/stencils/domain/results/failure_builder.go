package results

import (
	"errors"
	"strconv"

	"steve.care/network/libraries/hash"
)

type failureBuilder struct {
	hashAdapter     hash.Adapter
	code            uint
	isRaisedInLayer bool
	pIndex          *uint
}

func createFailureBuilder(
	hashAdapter hash.Adapter,
) FailureBuilder {
	out := failureBuilder{
		hashAdapter:     hashAdapter,
		code:            0,
		isRaisedInLayer: false,
		pIndex:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *failureBuilder) Create() FailureBuilder {
	return createFailureBuilder(
		app.hashAdapter,
	)
}

// WithIndex adds an index to the builder
func (app *failureBuilder) WithIndex(index uint) FailureBuilder {
	app.pIndex = &index
	return app
}

// WithCode adds a code to the builder
func (app *failureBuilder) WithCode(code uint) FailureBuilder {
	app.code = code
	return app
}

// IsRaisedInLayer flags the builder as isRaisedInLayer
func (app *failureBuilder) IsRaisedInLayer() FailureBuilder {
	app.isRaisedInLayer = true
	return app
}

// Now builds a new Failure instance
func (app *failureBuilder) Now() (Failure, error) {
	if app.code <= 0 {
		return nil, errors.New("the code is mandatory in order to build a Failure instance")
	}

	isRaisedInLayer := "false"
	if app.isRaisedInLayer {
		isRaisedInLayer = "true"
	}

	data := [][]byte{
		[]byte(strconv.Itoa(int(app.code))),
		[]byte(isRaisedInLayer),
	}

	if app.pIndex != nil {
		data = append(data, []byte(strconv.Itoa(int(*app.pIndex))))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pIndex != nil {
		return createFailureWithIndex(*pHash, app.code, app.isRaisedInLayer, app.pIndex), nil
	}

	return createFailure(*pHash, app.code, app.isRaisedInLayer), nil
}
