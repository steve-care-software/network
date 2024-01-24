package sqllites

import (
	"fmt"

	"steve.care/network/domain/queries/conditions"
)

func getSchema() string {
	return `
		DROP TABLE IF EXISTS accounts;
		CREATE TABLE accounts (
			username TEXT PRIMARY KEY,
			cipher BLOB
		);
	`
}

func processCondition(condition conditions.Condition, arguments []interface{}) (string, []interface{}) {
	pointer := condition.Pointer()
	queryPointer := pointerToString(pointer)

	operator := condition.Operator()
	queryOperator := operatorToField(operator)

	element := condition.Element()
	queryElement, retArguments := processElement(element, arguments)

	query := fmt.Sprintf("%s %s %s", queryPointer, queryOperator, queryElement)
	return query, retArguments
}

func processElement(element conditions.Element, arguments []interface{}) (string, []interface{}) {
	if element.IsCondition() {
		condition := element.Condition()
		queryCondition, retArguments := processCondition(condition, arguments)
		query := fmt.Sprintf("(%s)", queryCondition)
		return query, retArguments
	}

	criteria := element.Resource()
	return processResource(criteria, arguments)
}

func processResource(criteria conditions.Resource, arguments []interface{}) (string, []interface{}) {
	if criteria.IsField() {
		field := criteria.Field()
		return pointerToString(field), arguments
	}

	retArguments := append(arguments, criteria.Value())
	return "?", retArguments
}

func pointerToString(pointer conditions.Pointer) string {
	entity := pointer.Entity()
	field := pointer.Field()
	return fmt.Sprintf("%s.%s", entity, field)
}

func operatorToField(operator conditions.Operator) string {
	if operator.IsRelational() {
		relational := operator.Relational()
		return relationalOperatorToField(relational)
	}

	if operator.IsInteger() {
		integer := operator.Integer()
		return integerOperatorToField(integer)
	}

	return "="
}

func relationalOperatorToField(operator conditions.RelationalOperator) string {
	if operator.IsAnd() {
		return "&&"
	}

	return "||"
}

func integerOperatorToField(operator conditions.IntegerOperator) string {
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
