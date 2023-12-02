package transactions

// Transaction represents a transaction
type Transaction interface {
	Execute(query string, args ...any) error
	Rollback() error
	Commit() error
	Cancel() error
}
