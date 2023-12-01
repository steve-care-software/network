package sqllites

import (
	"database/sql"

	"steve.care/network/domain/databases"
)

type database struct {
	query databases.Query
	dbPtr *sql.DB
}

func createDatabase(
	query databases.Query,
	dbPtr *sql.DB,
) databases.Database {
	out := database{
		query: query,
		dbPtr: dbPtr,
	}

	return &out
}

// Execute executes a script in database
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

// Query returns the query
func (app *database) Query() databases.Query {
	return app.query
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
