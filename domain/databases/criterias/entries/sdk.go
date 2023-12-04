package entries

import (
	"steve.care/network/domain/databases/criterias/entries/resources"
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
	Now() (Entry, error)
}

// Entry represents a request entry
type Entry interface {
	Resource() resources.Resource
	Fields() []string
}
