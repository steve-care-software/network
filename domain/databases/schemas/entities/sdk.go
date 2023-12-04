package entities

import "steve.care/network/domain/databases/schemas/entities/fields"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewEntityBuilder creates a new entity builder
func NewEntityBuilder() EntityBuilder {
	return createEntityBuilder()
}

// Builder represents the entities builder
type Builder interface {
	Create() Builder
	WithList(list []Entity) Builder
	Now() (Entities, error)
}

// Entities represents entities
type Entities interface {
	List() []Entity
}

// EntityBuilder represents an entity builder
type EntityBuilder interface {
	Create() EntityBuilder
	WithName(name string) EntityBuilder
	WithHead(head string) EntityBuilder
	WithFields(fields fields.Fields) EntityBuilder
	Now() (Entity, error)
}

// Entity represents an entity
type Entity interface {
	Name() string
	Head() string
	Fields() fields.Fields
}
