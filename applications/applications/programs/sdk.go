package programs

import (
	"steve.care/network/domain/credentials"
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
	Execute(input []byte, program programs.Program, receipt receipts.Receipt) (receipts.Receipt, error)
}
