package programs

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs"
	"steve.care/network/domain/programs/blocks/executions"
	"steve.care/network/domain/receipts"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Root returns the root program
func (app *application) Root() (programs.Program, error) {
	return nil, nil
}

// Children returns the children paths
func (app *application) Children(path []string) ([]string, error) {
	return nil, nil
}

// Height returns the height
func (app *application) Height(path []string) (*uint, error) {
	return nil, nil
}

// Revision returns the revision
func (app *application) Revision(path []string, height uint) (hash.Hash, error) {
	return nil, nil
}

// Retrieve retrieves a program by hash
func (app *application) Retrieve(program hash.Hash) (programs.Program, error) {
	return nil, nil
}

// Insert inserts a program
func (app *application) Insert(path []string, description string) error {
	return nil
}

// Update updates a program
func (app *application) Update(program programs.Program, execution executions.Execution) error {
	return nil
}

// Convert converts a receipt to an execution
func (app *application) Convert(receipt receipts.Receipt) (executions.Execution, error) {
	return nil, nil
}

// Rewind rewinds a program
func (app *application) Rewind(path []string) error {
	return nil
}

// Delete deletes a program
func (app *application) Delete(path []string) error {
	return nil
}

// Execute executes a program
func (app *application) Execute(input []byte, program programs.Program, context receipts.Receipt) (receipts.Receipt, error) {
	return nil, nil
}
