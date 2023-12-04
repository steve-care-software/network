package entities

import "steve.care/network/domain/schemas/entities/fields"

// Entities represents entities
type Entities interface {
	List() []Entity
}

// Entity represents an entity
type Entity interface {
	Name() string
	Head() string
	Fields() fields.Fields
}
