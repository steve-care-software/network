package queries

import (
	"steve.care/network/domain/databases/criterias/entries"
	"steve.care/network/domain/hash"
)

// QueryFn represents the query fn
type QueryFn func(scannable Scannable) (interface{}, error)

// Scannable represents the scannable interface
type Scannable interface {
	Scan(dest ...any) error
}

// Query represents a query
type Query interface {
	Amount(container string) (uint, error)
	List(container string, index uint, amount uint) ([]hash.Hash, error)
	Retrieve(query entries.Entry) (interface{}, error)
	RetrieveByHash(hash hash.Hash) (interface{}, error)
	RetrieveList(container string, hashes []hash.Hash) ([]interface{}, error)
	QueryFirst(callback QueryFn, query string, args ...any) (interface{}, error)
	Query(callback QueryFn, query string, args ...any) ([]interface{}, error)
}
