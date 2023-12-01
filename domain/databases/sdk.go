package databases

// QueryFn represents the query fn
type QueryFn func(scannable Scannable) (interface{}, error)

// Scannable represents the scannable interface
type Scannable interface {
	Scan(dest ...any) error
}

// Database represents a database
type Database interface {
	Execute(script string) error
	Prepare() (Transaction, error)
	QueryFirst(callback QueryFn, query string, args ...any) (interface{}, error)
	Query(callback QueryFn, query string, args ...any) ([]interface{}, error)
	Close() error
}

// Transaction represents a transaction
type Transaction interface {
	Execute(query string, args ...any) (int64, error)
	Rollback() error
	Commit() error
	Cancel() error
}
