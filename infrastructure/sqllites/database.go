package sqllites

import (
	"database/sql"
	"errors"

	"steve.care/network/domain/databases"
)

type database struct {
	dbPtr *sql.DB
}

func createDatabase(
	dbPtr *sql.DB,
) databases.Database {
	out := database{
		dbPtr: dbPtr,
	}

	return &out
}

func (app *database) Execute(script string) error {
	_, err := app.dbPtr.Exec(script)
	if err != nil {
		return err
	}

	return nil
}

// Prepare prepares a transaction
func (app *database) Prepare() (databases.Transaction, error) {
	txPtr, err := app.dbPtr.Begin()
	if err != nil {
		return nil, err
	}

	return createTransaction(
		txPtr,
	), nil
}

// QueryFirst returns the first instance of a query
func (app *database) QueryFirst(callback databases.QueryFn, query string, args ...any) (interface{}, error) {
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
func (app *database) Query(callback databases.QueryFn, query string, args ...any) ([]interface{}, error) {
	return app.query(-1, callback, query, args...)
}

func (app *database) query(max int, callback databases.QueryFn, query string, args ...any) ([]interface{}, error) {
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

// Close closes the database
func (app *database) Close() error {
	err := app.dbPtr.Close()
	if err != nil {
		return err
	}

	app.dbPtr = nil
	return nil
}
