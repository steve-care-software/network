package schemas

import (
	"steve.care/network/domain/databases/schemas/connections"
	"steve.care/network/domain/databases/schemas/entities"
)

type schema struct {
	entities    entities.Entities
	connections connections.Connections
}

func createSchema(
	entities entities.Entities,
) Schema {
	return createSchemaInternally(
		entities,
		nil,
	)
}

func createSchemaWithConnections(
	entities entities.Entities,
	connections connections.Connections,
) Schema {
	return createSchemaInternally(
		entities,
		connections,
	)
}

func createSchemaInternally(
	entities entities.Entities,
	connections connections.Connections,
) Schema {
	out := schema{
		entities:    entities,
		connections: connections,
	}

	return &out
}

// Entities return the entities
func (obj *schema) Entities() entities.Entities {
	return obj.entities
}

// HasConnections returns true if there is connections, false otherwise
func (obj *schema) HasConnections() bool {
	return obj.connections != nil
}

// Connections returns the connections, if any
func (obj *schema) Connections() connections.Connections {
	return obj.connections
}
