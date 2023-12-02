package queries

// QueryFn represents the query fn
type QueryFn func(scannable Scannable) (interface{}, error)

// Scannable represents the scannable interface
type Scannable interface {
	Scan(dest ...any) error
}

// Query represents a query
type Query interface {
	QueryFirst(callback QueryFn, query string, args ...any) (interface{}, error)
	Query(callback QueryFn, query string, args ...any) ([]interface{}, error)
}
