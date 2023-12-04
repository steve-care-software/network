package transactions

import (
	"steve.care/network/domain/databases/criterias/entries/resources"
)

// Transaction represents a transaction
type Transaction interface {
	Insert(container string, values map[string]interface{}) error
	Update(original resources.Resource, updatedValues map[string]interface{}) error
	Delete(resource resources.Resource) error
	Execute(query string, args ...any) error
	Rollback() error
	Commit() error
	Cancel() error
}
