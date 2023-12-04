package connections

import "steve.care/network/domain/databases/schemas/connections/pointers"

// Builder represents the connections builder
type Builder interface {
	Create() Builder
	WithList(list []Connection) Builder
	Now() (Connections, error)
}

// Connections represents connections
type Connections interface {
	List() []Connection
}

// ConnectionBuilder represents a connection builder
type ConnectionBuilder interface {
	Create() ConnectionBuilder
	From(from pointers.Pointer) ConnectionBuilder
	To(to pointers.Pointer) ConnectionBuilder
	Now() (Connection, error)
}

// Connection represents a connection
type Connection interface {
	From() pointers.Pointer
	To() pointers.Pointer
}
