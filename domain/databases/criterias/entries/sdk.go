package entries

import (
	"steve.care/network/domain/databases/criterias"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithCriteria(criteria criterias.Criteria) Builder
	WithFields(fields []string) Builder
	Now() (Entry, error)
}

// Entry represents a request entry
type Entry interface {
	Criteria() criterias.Criteria
	Fields() []string
}
