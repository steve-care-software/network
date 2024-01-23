package daemons

import (
	"steve.care/network/applications/applications/programs/graphs"
	"steve.care/network/applications/applications/programs/logics"
	"steve.care/network/applications/applications/programs/peers"
)

// NewProgram creates a new program application
func NewProgram(
	graphApplication graphs.Application,
	peersApplication peers.Application,
	logicApplication logics.Application,
) Application {
	return createApplication(
		graphApplication,
		peersApplication,
		logicApplication,
	)
}

// Application represents a deamon application
type Application interface {
	Start() error
	Stop()
}
