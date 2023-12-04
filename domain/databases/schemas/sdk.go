package schemas

import (
	"steve.care/network/domain/databases/schemas/connections"
	"steve.care/network/domain/databases/schemas/entities"
)

// Builder represents the schema builder
type Builder interface {
	Create() Builder
	WithEntities(entities entities.Entities) Builder
	WithConnections(connections connections.Connections) Builder
	Now() (Schema, error)
}

// Schema represents a schame
type Schema interface {
	Entities() entities.Entities
	HasConnections() bool
	Connections() connections.Connections
}
