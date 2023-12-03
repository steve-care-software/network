package entities

import (
	"steve.care/network/domain/databases/criterias/entities/resources"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithResource(resource resources.Resource) Builder
	WithFields(fields []string) Builder
	Now() (Entity, error)
}

// Entity represents a request entity
type Entity interface {
	Resource() resources.Resource
	Fields() []string
}
