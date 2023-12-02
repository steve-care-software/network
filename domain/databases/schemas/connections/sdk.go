package connections

import "steve.care/network/domain/databases/schemas/connections/pointers"

// Connections represents connections
type Connections interface {
	List() []Connection
}

// Connection represents a container connection
type Connection interface {
	From() pointers.Pointer
	To() pointers.Pointer
}
