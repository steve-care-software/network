package schemas

import (
	"steve.care/network/domain/databases/schemas/connections"
	"steve.care/network/domain/databases/schemas/containers"
)

// Schema represents a schema
type Schema interface {
	Containers() containers.Containers
	HasConnections() bool
	Connections() connections.Connections
}
