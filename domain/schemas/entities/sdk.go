package entities

import "steve.care/network/domain/schemas/entities/fields"

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
	WithNote(note string) EntityBuilder
	Now() (Entity, error)
}

// Entity represents an entity
type Entity interface {
	Name() string
	Head() string
	Fields() fields.Fields
	HasNote() bool
	Note() string
}
