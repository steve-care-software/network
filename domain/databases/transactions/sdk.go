package transactions

import (
	"steve.care/network/domain/databases/criterias"
)

// Transaction represents a transaction
type Transaction interface {
	Insert(container string, values map[string]interface{}) error
	Update(original criterias.Criteria, updatedValues map[string]interface{}) error
	Delete(criteria criterias.Criteria) error
	Execute(query string, args ...any) error
	Rollback() error
	Commit() error
	Cancel() error
}
