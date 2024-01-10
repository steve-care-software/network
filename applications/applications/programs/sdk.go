package programs

import (
	"steve.care/network/applications/applications/programs/blocks"
	"steve.care/network/applications/applications/programs/cruds"
	"steve.care/network/applications/applications/programs/logics"
	"steve.care/network/applications/applications/programs/peers"
	"steve.care/network/applications/applications/programs/threads"
	"steve.care/network/domain/credentials"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the program application
type Application interface {
	Block() blocks.Application
	CRUD() cruds.Application
	Logic() logics.Application
	Threads() threads.Application
	Peers() peers.Application
}
