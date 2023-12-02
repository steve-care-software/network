package pointers

import "steve.care/network/domain/databases/schemas/containers"

// Pointer represent sa connection pointer
type Pointer interface {
	Container() containers.Container
	Field() string
}
