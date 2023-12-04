package connections

import "steve.care/network/domain/schemas/connections/pointers"

// Connections represents connections
type Connections interface {
	List() []Connection
}

// Connection represents a connection
type Connection interface {
	From() pointers.Pointer
	To() pointers.Pointer
}
