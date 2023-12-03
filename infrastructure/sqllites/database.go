package sqllites

import (
	"database/sql"

	"steve.care/network/domain/databases"
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/databases/transactions"
)

type database struct {
	query queries.Query
	dbPtr *sql.DB
}

func createDatabase(
	query queries.Query,
	dbPtr *sql.DB,
) databases.Database {
	out := database{
		query: query,
		dbPtr: dbPtr,
	}

	return &out
}

// Execute executes a schema on the database
func (app *database) Execute(schema string) error {
	_, err := app.dbPtr.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

// Prepare prepares a transaction
func (app *database) Prepare() (transactions.Transaction, error) {
	txPtr, err := app.dbPtr.Begin()
	if err != nil {
		return nil, err
	}

	return createTransaction(
		txPtr,
	), nil
}

// Query returns the query
func (app *database) Query() queries.Query {
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
