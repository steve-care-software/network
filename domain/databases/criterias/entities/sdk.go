package entities

import (
	"steve.care/network/domain/databases/criterias/entities/resources"
)

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithResource(resource resources.Resource)
	WithFields(fields []string) Builder
	Now() (Entity, error)
}

// Entity represents a request entity
type Entity interface {
	Resource() resources.Resource
	Fields() []string
}
