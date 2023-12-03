package sqllites

import (
	"database/sql"
	"fmt"
	"strings"

	"steve.care/network/domain/databases/criterias/entities/resources"
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
func (app *transaction) Insert(container string, values map[string]interface{}) error {
	fieldValuesList := []any{}
	fieldValuePlaceholders := []string{}
	fieldNamesList := []string{}
	for keyname, oneValue := range values {
		fieldValuesList = append(fieldValuesList, oneValue)
		fieldValuePlaceholders = append(fieldValuePlaceholders, "?")
		fieldNamesList = append(fieldNamesList, keyname)
	}

	fieldValuesStr := strings.Join(fieldNamesList, ", ")
	fieldValuePlaceholdersStr := strings.Join(fieldValuePlaceholders, ", ")
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", container, fieldValuesStr, fieldValuePlaceholdersStr)
	_, err := app.txPtr.Exec(query, fieldValuesList...)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a resource
func (app *transaction) Update(original resources.Resource, updatedValues map[string]interface{}) error {
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
