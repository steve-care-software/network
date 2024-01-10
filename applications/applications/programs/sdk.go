package programs

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs"
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
	List(isActive bool) ([]hash.Hash, error)
	//Root() (programs.Program, error)
	//Children(path []string) ([]string, error)
	//Height(path []string) (*uint, error)
	//Revision(path []string, height uint) (hash.Hash, error)
	Retrieve(program hash.Hash) (programs.Program, error)
	//Insert(path []string, description string) error
	//Update(program programs.Program, trx transactions.Transactions) error
	//Convert(receipt receipts.Receipt) (executions.Execution, error)
	//Rewind(path []string) error
	//Delete(path []string) error
	Execute(input []byte, program programs.Program, context receipts.Receipt) (receipts.Receipt, error)
}
