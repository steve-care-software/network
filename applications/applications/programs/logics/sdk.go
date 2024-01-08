package logics

import (
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/programs/logics/libraries"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/receipts"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the logic application
type Application interface {
	Execute(input []byte, layer layers.Layer, library libraries.Library, context receipts.Receipt) (receipts.Receipt, error)
}
