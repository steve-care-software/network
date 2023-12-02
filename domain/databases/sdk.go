package databases

import (
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/databases/transactions"
)

// Database represents a database
type Database interface {
	Execute(script string) error
	Prepare() (transactions.Transaction, error)
	Query() queries.Query
	Close() error
}
