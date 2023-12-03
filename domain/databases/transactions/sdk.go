package transactions

import (
	"steve.care/network/domain/databases/criterias/entities/resources"
	"steve.care/network/domain/databases/criterias/values"
)

// Transaction represents a transaction
type Transaction interface {
	Insert(container string, values map[string]values.Value) error
	Update(original resources.Resource, updatedValues map[string]values.Value) error
	Delete(resource resources.Resource) error
	Execute(query string, args ...any) error
	Rollback() error
	Commit() error
	Cancel() error
}
