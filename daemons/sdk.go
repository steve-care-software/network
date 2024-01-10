package daemons

import (
	"steve.care/network/applications/applications/programs/cruds"
	"steve.care/network/applications/applications/programs/logics"
	"steve.care/network/applications/applications/programs/peers"
)

// NewProgram creates a new program application
func NewProgram(
	crudApplication cruds.Application,
	peersApplication peers.Application,
	logicApplication logics.Application,
) Application {
	return createApplication(
		crudApplication,
		peersApplication,
		logicApplication,
	)
}

// Application represents a deamon application
type Application interface {
	Start() error
	Stop()
}
