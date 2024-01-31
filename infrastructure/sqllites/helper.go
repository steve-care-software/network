package sqllites

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

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

func callMethodsOnInstance(
	methods []string,
	pInstance interface{},
	pErrorStr *string,
) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			value := fmt.Sprint(r)
			*pErrorStr = value
		}
	}()

	value := reflect.ValueOf(pInstance)
	for _, oneMethod := range methods {
		if value.IsNil() {
			return nil, nil
		}

		retValues := value.MethodByName(oneMethod).Call([]reflect.Value{})
		if len(retValues) < 1 {
			str := fmt.Sprintf("at least %d values were returned, %d were expected, when calling the method (name %s) in the method chain (%s)", len(retValues), 1, oneMethod, strings.Join(methods, ","))
			return nil, errors.New(str)
		}

		value = retValues[0]
	}

	return value.Interface(), nil
}

func callMethodOnInstanceWithParams(
	method string,
	pInstance interface{},
	pErrorStr *string,
	params []interface{},
) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			value := fmt.Sprint(r)
			*pErrorStr = value
		}
	}()

	value := reflect.ValueOf(pInstance)
	methodName := value.MethodByName(method)
	if !methodName.IsValid() {
		str := fmt.Sprintf("there is no method (name: %s) on the provided instance", method)
		return nil, errors.New(str)
	}

	methodAmountArguments := methodName.Type().NumIn()
	if methodAmountArguments != len(params) {
		str := fmt.Sprintf("the methodName (%s) was expected to contain %d arguments, but it contains %d arguments in reality", method, len(params), methodAmountArguments)
		return nil, errors.New(str)
	}

	methodAmountReturns := methodName.Type().NumIn()
	if methodName.Type().NumOut() != 1 {
		str := fmt.Sprintf("the methodName (%s) was expected to contain %d retrun values, but it contains %d arguments in reality", method, 1, methodAmountReturns)
		return nil, errors.New(str)
	}

	methodParams := []reflect.Value{}
	if params != nil && len(params) > 0 {
		for _, oneParam := range params {

			expectedType := methodName.Type().In(0)
			value := reflect.ValueOf(oneParam)
			currentType := value.Type()

			// if the types are different, try to conver it:
			if expectedType.Kind() != currentType.Kind() {
				if value.CanConvert(expectedType) {
					value = value.Convert(expectedType)
				}
			}

			methodParams = append(methodParams, value)
		}
	}

	retValues := value.MethodByName(method).Call(methodParams)
	if len(retValues) < 1 {
		str := fmt.Sprintf("%d  values were returned, at least %d were expected, when calling the method (name %s)", len(retValues), 1, method)
		return nil, errors.New(str)
	}

	return retValues[0].Interface(), nil
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
