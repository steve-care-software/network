package sqllites

import (
	"database/sql"

	"steve.care/network/domain/databases/criterias/entities/resources"
	"steve.care/network/domain/databases/criterias/values"
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

// Insert inserts a resource
func (app *transaction) Insert(container string, values map[string]values.Value) error {
	return nil
}

// Update updates a resource
func (app *transaction) Update(original resources.Resource, updatedValues map[string]values.Value) error {
	return nil
}

// Delete deletes a resource
func (app *transaction) Delete(resource resources.Resource) error {
	return nil
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
