package executions

import (
	"steve.care/network/domain/programs"
	"steve.care/network/domain/receipts"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Execute executes an application
func (app *application) Execute(input []byte, program programs.Program, context receipts.Receipt) (receipts.Receipt, error) {
	// execute the program:

	// build the new origin:

	// retrieve the link related to our new origin, if any:

	// execute the link:
	return nil, nil
}
