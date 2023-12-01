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
	Query() Query
	Close() error
}

// Query represents a query
type Query interface {
	QueryFirst(callback QueryFn, query string, args ...any) (interface{}, error)
	Query(callback QueryFn, query string, args ...any) ([]interface{}, error)
}

// Transaction represents a transaction
type Transaction interface {
	Execute(query string, args ...any) error
	Rollback() error
	Commit() error
	Cancel() error
}
