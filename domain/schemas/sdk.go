package schemas

import (
	"steve.care/network/domain/schemas/connections"
	"steve.care/network/domain/schemas/entities"
)

// Schema represents a schame
type Schema interface {
	Entities() entities.Entities
	HasConnections() bool
	Connections() connections.Connections
}
