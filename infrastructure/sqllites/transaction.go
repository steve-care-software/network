package sqllites

import (
	"database/sql"
	"fmt"
	"strings"

	"steve.care/network/domain/databases/criterias"
	"steve.care/network/domain/databases/criterias/conditions"
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

// Insert inserts a criteria
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

// Update updates a criteria
func (app *transaction) Update(original criterias.Criteria, updatedValues map[string]interface{}) error {
	return nil
}

// Delete deletes a criteria
func (app *transaction) Delete(criteria criterias.Criteria) error {
	entity := criteria.Entity()
	condition := criteria.Condition()
	whereClause, arguments := app.processCondition(condition, []interface{}{})
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", entity, whereClause)
	_, err := app.txPtr.Exec(query, arguments...)
	if err != nil {
		return err
	}

	return nil
}

func (app *transaction) processCondition(condition conditions.Condition, arguments []interface{}) (string, []interface{}) {
	pointer := condition.Pointer()
	queryPointer := app.pointerToString(pointer)

	operator := condition.Operator()
	queryOperator := app.operatorToField(operator)

	element := condition.Element()
	queryElement, retArguments := app.processElement(element, arguments)

	query := fmt.Sprintf("%s %s %s", queryPointer, queryOperator, queryElement)
	return query, retArguments
}

func (app *transaction) processElement(element conditions.Element, arguments []interface{}) (string, []interface{}) {
	if element.IsCondition() {
		condition := element.Condition()
		queryCondition, retArguments := app.processCondition(condition, arguments)
		query := fmt.Sprintf("(%s)", queryCondition)
		return query, retArguments
	}

	criteria := element.Resource()
	return app.processResource(criteria, arguments)
}

func (app *transaction) processResource(criteria conditions.Resource, arguments []interface{}) (string, []interface{}) {
	if criteria.IsField() {
		field := criteria.Field()
		return app.pointerToString(field), arguments
	}

	retArguments := append(arguments, criteria.Value())
	return "?", retArguments
}

func (app *transaction) pointerToString(pointer conditions.Pointer) string {
	entity := pointer.Entity()
	field := pointer.Field()
	return fmt.Sprintf("%s.%s", entity, field)
}

func (app *transaction) operatorToField(operator conditions.Operator) string {
	if operator.IsRelational() {
		relational := operator.Relational()
		return app.relationalOperatorToField(relational)
	}

	if operator.IsInteger() {
		integer := operator.Integer()
		return app.integerOperatorToField(integer)
	}

	return "="
}

func (app *transaction) relationalOperatorToField(operator conditions.RelationalOperator) string {
	if operator.IsAnd() {
		return "&&"
	}

	return "||"
}

func (app *transaction) integerOperatorToField(operator conditions.IntegerOperator) string {
	if operator.IsSmallerThan() && operator.IsEqual() {
		return "<="
	}

	if operator.IsSmallerThan() {
		return "<"
	}

	if operator.IsBiggerThan() && operator.IsEqual() {
		return ">="
	}

	if operator.IsBiggerThan() {
		return ">"
	}

	if operator.IsEqual() {
		return "="
	}

	return "!="
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
