package resources

import (
	"steve.care/network/domain/schemas/groups/resources/fields"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewResourceBuilder creates a new resources builder
func NewResourceBuilder() ResourceBuilder {
	return createResourceBuilder()
}

// NewConnectionsBuilder creates a new connections builder
func NewConnectionsBuilder() ConnectionsBuilder {
	return createConnectionsBuilder()
}

// NewConnectionBuilder creates a new connection builder
func NewConnectionBuilder() ConnectionBuilder {
	return createConnectionBuilder()
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	return createPointerBuilder()
}

// Builder represents the resources builder
type Builder interface {
	Create() Builder
	WithList(list []Resource) Builder
	Now() (Resources, error)
}

// Resources represents a resource
type Resources interface {
	List() []Resource
	Fetch(name string) (Resource, error)
}

// ResourceBuilder represents a resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithName(name string) ResourceBuilder
	WithKey(key fields.Field) ResourceBuilder
	WithFields(fields fields.Fields) ResourceBuilder
	WithConnections(connections Connections) ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a schema resource
type Resource interface {
	Name() string
	Key() fields.Field
	Fields() fields.Fields
	HasConnections() bool
	Connections() Connections
}

// ConnectionsBuilder represents a connections builder
type ConnectionsBuilder interface {
	Create() ConnectionsBuilder
	WithList(list []Connection) ConnectionsBuilder
	Now() (Connections, error)
}

// Connections represents connections
type Connections interface {
	List() []Connection
}

// ConnectionBuilder represents a connection builder
type ConnectionBuilder interface {
	Create() ConnectionBuilder
	WithField(field string) ConnectionBuilder
	WithReference(reference Pointer) ConnectionBuilder
	Now() (Connection, error)
}

// Connection represenst a connection
type Connection interface {
	Field() string
	Reference() Pointer
}

// PointerBuilder represents a pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithResource(resource Resource) PointerBuilder
	WithField(field string) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a connection pointer
type Pointer interface {
	Resource() Resource
	Field() string
}
