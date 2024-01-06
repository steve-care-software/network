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
	Children(path []hash.Hash) ([]hash.Hash, error)
	Retrieve(path []hash.Hash) (programs.Program, error)
	Insert(name string, description string) error
	Update(program programs.Program, execution executions.Execution) error
	Convert(receipt receipts.Receipt) (executions.Execution, error)
	Delete(hash hash.Hash) error
	Execute(input []byte, program programs.Program, context receipts.Receipt) (receipts.Receipt, error)
}
