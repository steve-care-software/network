package sqllites

import (
	"database/sql"
	"errors"

	"steve.care/network/domain/databases/criterias/entities"
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/hash"
)

type query struct {
	dbPtr *sql.DB
}

func createQuery(
	dbPtr *sql.DB,
) queries.Query {
	out := query{
		dbPtr: dbPtr,
	}

	return &out
}

// Amount returns the amount of entities a container contains
func (app *query) Amount(container string) (uint, error) {
	return 0, nil
}

// List lists entity hashes
func (app *query) List(container string, index uint, amount uint) ([]hash.Hash, error) {
	return nil, nil
}

// Retrieve retrieves an entity by query
func (app *query) Retrieve(query entities.Entity) (interface{}, error) {
	return nil, nil
}

// RetrieveByHash retrieves an entity by hash
func (app *query) RetrieveByHash(hash hash.Hash) (interface{}, error) {
	return nil, nil
}

// RetrieveList retrieves a list of entity by hashes
func (app *query) RetrieveList(container string, hashes []hash.Hash) ([]interface{}, error) {
	return nil, nil
}

// QueryFirst returns the first instance of a query
func (app *query) QueryFirst(callback queries.QueryFn, query string, args ...any) (interface{}, error) {
	list, err := app.query(-1, callback, query, args...)
	if err != nil {
		return nil, err
	}

	if len(list) <= 0 {
		return nil, errors.New("there was no element returned from the query")
	}

	return list[0], nil
}

// Query executes a query
func (app *query) Query(callback queries.QueryFn, query string, args ...any) ([]interface{}, error) {
	return app.query(-1, callback, query, args...)
}

func (app *query) query(max int, callback queries.QueryFn, query string, args ...any) ([]interface{}, error) {
	rows, err := app.dbPtr.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	output := []interface{}{}
	for rows.Next() {
		ins, err := callback(rows)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
		if max < 0 {
			continue
		}

		if len(output) <= max {
			break
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return output, nil
}
