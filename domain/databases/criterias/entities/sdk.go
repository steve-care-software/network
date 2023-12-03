package entities

import "steve.care/network/domain/databases/criterias/conditions"

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithContainer(container string) Builder
	WithFields(fields []string) Builder
	WithCondition(condition conditions.Condition) Builder
	Now() (Entity, error)
}

// Entity represents a request entity
type Entity interface {
	Container() string
	Fields() []string
	Condition() conditions.Condition
}
