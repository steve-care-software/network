package logics

import (
	"steve.care/network/domain/programs/logics/libraries"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/receipts"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Execute executes the logic application
func (app *application) Execute(input []byte, layer layers.Layer, library libraries.Library, context receipts.Receipt) (receipts.Receipt, error) {
	// execute the layer:

	// build the new origin:

	// retrieve the link related to our new origin, if any, from our library:

	// execute the link:
	return nil, nil
}
