package containers

import "steve.care/network/domain/databases/schemas/containers/fields"

// Containers represents containers
type Containers interface {
	List() []Container
}

// Container represents a container
type Container interface {
	Name() string
	Head() fields.Field
	Fields() fields.Fields
}
