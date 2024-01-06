package programs

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs"
	"steve.care/network/domain/programs/blocks/executions"
	"steve.care/network/domain/receipts"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the program application
type Application interface {
	Root() (programs.Program, error)
	Children(spacesPath [][]string) ([][]string, error)
	Revisions(spacesPath [][]string) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (programs.Program, error)
	Insert(space []string, description string) error
	Update(program programs.Program, execution executions.Execution) error
	Convert(receipt receipts.Receipt) (executions.Execution, error)
	Rewind(space []string) error
	Delete(space []string) error
	Execute(input []byte, program programs.Program, context receipts.Receipt) (receipts.Receipt, error)
}
