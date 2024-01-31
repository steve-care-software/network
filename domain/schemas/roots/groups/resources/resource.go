package resources

import (
	"steve.care/network/domain/schemas/roots/groups/resources/fields"
	"steve.care/network/domain/schemas/roots/groups/methods"
)

type resource struct {
	name        string
	key         fields.Field
	fields      fields.Fields
	builder     methods.Methods
	connections Connections
}

func createResource(
	name string,
	key fields.Field,
	fields fields.Fields,
	builder methods.Methods,
) Resource {
	return createResourceInternally(name, key, fields, builder, nil)
}

func createResourceWithConnections(
	name string,
	key fields.Field,
	fields fields.Fields,
	builder methods.Methods,
	connections Connections,
) Resource {
	return createResourceInternally(name, key, fields, builder, connections)
}

func createResourceInternally(
	name string,
	key fields.Field,
	fields fields.Fields,
	builder methods.Methods,
	connections Connections,
) Resource {
	out := resource{
		name:        name,
		key:         key,
		fields:      fields,
		builder:     builder,
		connections: connections,
	}

	return &out
}

// Name returns the name
func (obj *resource) Name() string {
	return obj.name
}

// Key returns the key
func (obj *resource) Key() fields.Field {
	return obj.key
}

// Fields returns the fields
func (obj *resource) Fields() fields.Fields {
	return obj.fields
}

// Builder returns the builder methods
func (obj *resource) Builder() methods.Methods {
	return obj.builder
}

// HasConnections returns true if there is connections, false otherwise
func (obj *resource) HasConnections() bool {
	return obj.connections != nil
}

// Connections returns the connections, if any
func (obj *resource) Connections() Connections {
	return obj.connections
}
