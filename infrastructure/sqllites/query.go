package sqllites

import (
	"database/sql"
	"errors"

	"steve.care/network/domain/databases/queries"
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
