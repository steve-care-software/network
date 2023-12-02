package inserts

import "steve.care/network/domain/databases/values"

// Builder represents an insert
type Builder interface {
}

// Insert represents an insert
type Insert interface {
	Container() string
	Values() map[string]values.Value
}
