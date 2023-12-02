package sqllites

import (
	"database/sql"

	"steve.care/network/domain/databases/transactions"
)

type transaction struct {
	txPtr *sql.Tx
}

func createTransaction(
	txPtr *sql.Tx,
) transactions.Transaction {
	out := transaction{
		txPtr: txPtr,
	}

	return &out
}

// Execute executes a transactional query
func (app *transaction) Execute(query string, args ...any) error {
	_, err := app.txPtr.Exec(query, args...)
	if err != nil {
		return nil
	}

	return nil
}

// Rollback the transaction
func (app *transaction) Rollback() error {
	return app.txPtr.Rollback()
}

// Commit commits the transaction
func (app *transaction) Commit() error {
	return app.txPtr.Commit()
}

// Cancel cancels the transaction
func (app *transaction) Cancel() error {
	err := app.txPtr.Rollback()
	if err != nil {
		return err
	}

	app.txPtr = nil
	return nil
}
