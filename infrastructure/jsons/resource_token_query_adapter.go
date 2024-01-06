package jsons

import (
	"steve.care/network/domain/queries"
	"steve.care/network/domain/queries/conditions"

	resources_queries "steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/queries"
	structs_tokens "steve.care/network/infrastructure/jsons/resources/tokens"
	structs_queries "steve.care/network/infrastructure/jsons/resources/tokens/queries"
)

type resourceTokenQueryAdapter struct {
	builder                   resources_queries.Builder
	queryBuilder              queries.Builder
	conditionBuilder          conditions.Builder
	pointerBulder             conditions.PointerBuilder
	elementBuilder            conditions.ElementBuilder
	resourceBuilder           conditions.ResourceBuilder
	operatorBuilder           conditions.OperatorBuilder
	relationalOperatorBuilder conditions.RelationalOperatorBuilder
	integerOperatorBuilder    conditions.IntegerOperatorBuilder
}

func (app *resourceTokenQueryAdapter) toStruct(ins resources_queries.Query) structs_tokens.Query {
	output := structs_tokens.Query{}
	if ins.IsQuery() {
		query := app.queryToStruct(ins.Query())
		output.Query = &query
	}

	if ins.IsCondition() {
		condition := app.conditionToStruct(ins.Condition())
		output.Condition = &condition
	}

	if ins.IsPointer() {
		pointer := app.pointerToStruct(ins.Pointer())
		output.Pointer = &pointer
	}

	if ins.IsElement() {
		element := app.elementToStruct(ins.Element())
		output.Element = &element
	}

	if ins.IsResource() {
		resource := app.resourceToStruct(ins.Resource())
		output.Resource = &resource
	}

	if ins.IsOperator() {
		operator := app.operatorToStruct(ins.Operator())
		output.Operator = &operator
	}

	if ins.IsRelationalOperator() {
		relationalOperator := app.relationalOperatorToStruct(ins.RelationalOperator())
		output.RelationalOperator = &relationalOperator
	}

	if ins.IsIntegerOperator() {
		integerOperator := app.integerOperatorToStruct(ins.IntegerOperator())
		output.IntegerOperator = &integerOperator
	}

	return output
}

func (app *resourceTokenQueryAdapter) toInstance(ins structs_tokens.Query) (resources_queries.Query, error) {
	builder := app.builder.Create()
	if ins.Query != nil {
		query, err := app.structToQuery(*ins.Query)
		if err != nil {
			return nil, err
		}

		builder.WithQuery(query)
	}

	if ins.Condition != nil {
		condition, err := app.structToCondtion(*ins.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	if ins.Pointer != nil {
		pointer, err := app.structToPointer(*ins.Pointer)
		if err != nil {
			return nil, err
		}

		builder.WithPointer(pointer)
	}

	if ins.Element != nil {
		element, err := app.structToElement(*ins.Element)
		if err != nil {
			return nil, err
		}

		builder.WithElement(element)
	}

	if ins.Resource != nil {
		resource, err := app.structToResource(*ins.Resource)
		if err != nil {
			return nil, err
		}

		builder.WithResource(resource)
	}

	if ins.Operator != nil {
		operator, err := app.structToOperator(*ins.Operator)
		if err != nil {
			return nil, err
		}

		builder.WithOperator(operator)
	}

	if ins.RelationalOperator != nil {
		relational, err := app.structToRelationalOperator(*ins.RelationalOperator)
		if err != nil {
			return nil, err
		}

		builder.WithRelationalOperator(relational)
	}

	if ins.IntegerOperator != nil {
		integer, err := app.structToIntegerOperator(*ins.IntegerOperator)
		if err != nil {
			return nil, err
		}

		builder.WithIntegerOperator(integer)
	}

	return builder.Now()
}

func (app *resourceTokenQueryAdapter) queryToStruct(
	ins queries.Query,
) structs_queries.Query {
	condition := app.conditionToStruct(ins.Condition())
	output := structs_queries.Query{
		Entity:    ins.Entity(),
		Condition: condition,
	}

	if ins.HasFields() {
		output.Fields = ins.Fields()
	}

	return output
}

func (app *resourceTokenQueryAdapter) structToQuery(
	ins structs_queries.Query,
) (queries.Query, error) {
	condition, err := app.structToCondtion(ins.Condition)
	if err != nil {
		return nil, err
	}

	builder := app.queryBuilder.Create().
		WithEntity(ins.Entity).
		WithCondition(condition)

	if ins.Fields != nil && len(ins.Fields) > 0 {
		builder.WithFields(ins.Fields)
	}

	return builder.Now()
}

func (app *resourceTokenQueryAdapter) conditionToStruct(
	ins conditions.Condition,
) structs_queries.Condition {
	pointer := app.pointerToStruct(ins.Pointer())
	operator := app.operatorToStruct(ins.Operator())
	element := app.elementToStruct(ins.Element())
	return structs_queries.Condition{
		Pointer:  pointer,
		Operator: operator,
		Element:  element,
	}
}

func (app *resourceTokenQueryAdapter) structToCondtion(
	ins structs_queries.Condition,
) (conditions.Condition, error) {
	pointer, err := app.structToPointer(ins.Pointer)
	if err != nil {
		return nil, err
	}

	operator, err := app.structToOperator(ins.Operator)
	if err != nil {
		return nil, err
	}

	element, err := app.structToElement(ins.Element)
	if err != nil {
		return nil, err
	}

	return app.conditionBuilder.Create().
		WithPointer(pointer).
		WithOperator(operator).
		WithElement(element).
		Now()
}

func (app *resourceTokenQueryAdapter) pointerToStruct(
	ins conditions.Pointer,
) structs_queries.Pointer {
	return structs_queries.Pointer{
		Entity: ins.Entity(),
		Field:  ins.Field(),
	}
}

func (app *resourceTokenQueryAdapter) structToPointer(
	ins structs_queries.Pointer,
) (conditions.Pointer, error) {
	return app.pointerBulder.Create().
		WithEntity(ins.Entity).
		WithField(ins.Field).
		Now()
}

func (app *resourceTokenQueryAdapter) elementToStruct(
	ins conditions.Element,
) structs_queries.Element {
	output := structs_queries.Element{}
	if ins.IsCondition() {
		condition := app.conditionToStruct(ins.Condition())
		output.Condition = &condition
	}

	if ins.IsResource() {
		resource := app.resourceToStruct(ins.Resource())
		output.Resource = &resource
	}

	return output
}

func (app *resourceTokenQueryAdapter) structToElement(
	ins structs_queries.Element,
) (conditions.Element, error) {
	builder := app.elementBuilder.Create()
	if ins.Condition != nil {
		condition, err := app.structToCondtion(*ins.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	if ins.Resource != nil {
		resource, err := app.structToResource(*ins.Resource)
		if err != nil {
			return nil, err
		}

		builder.WithResource(resource)
	}

	return builder.Now()
}

func (app *resourceTokenQueryAdapter) resourceToStruct(
	ins conditions.Resource,
) structs_queries.Resource {
	output := structs_queries.Resource{}
	if ins.IsField() {
		field := app.pointerToStruct(ins.Field())
		output.Field = &field
	}

	if ins.IsValue() {
		output.Value = ins.Value()
	}

	return output
}

func (app *resourceTokenQueryAdapter) structToResource(
	ins structs_queries.Resource,
) (conditions.Resource, error) {
	builder := app.resourceBuilder.Create()
	if ins.Field != nil {
		field, err := app.structToPointer(*ins.Field)
		if err != nil {
			return nil, err
		}

		builder.WithField(field)
	}

	if ins.Value != nil {
		builder.WithValue(ins.Value)
	}

	return builder.Now()
}

func (app *resourceTokenQueryAdapter) operatorToStruct(
	ins conditions.Operator,
) structs_queries.Operator {
	output := structs_queries.Operator{}
	if ins.IsEqual() {
		output.Equal = ins.IsEqual()
	}

	if ins.IsRelational() {
		relational := app.relationalOperatorToStruct(ins.Relational())
		output.Relational = &relational
	}

	if ins.IsInteger() {
		integer := app.integerOperatorToStruct(ins.Integer())
		output.Integer = &integer
	}

	return output
}

func (app *resourceTokenQueryAdapter) structToOperator(
	ins structs_queries.Operator,
) (conditions.Operator, error) {
	builder := app.operatorBuilder.Create()
	if ins.Equal {
		builder.IsEqual()
	}

	if ins.Relational != nil {
		relational, err := app.structToRelationalOperator(*ins.Relational)
		if err != nil {
			return nil, err
		}

		builder.WithRelational(relational)
	}

	if ins.Integer != nil {
		integer, err := app.structToIntegerOperator(*ins.Integer)
		if err != nil {
			return nil, err
		}

		builder.WithInteger(integer)
	}

	return builder.Now()
}

func (app *resourceTokenQueryAdapter) relationalOperatorToStruct(
	ins conditions.RelationalOperator,
) structs_queries.RelationalOperator {
	output := structs_queries.RelationalOperator{}
	if ins.IsAnd() {
		output.And = ins.IsAnd()
	}

	if ins.IsOr() {
		output.Or = ins.IsOr()
	}

	return output
}

func (app *resourceTokenQueryAdapter) structToRelationalOperator(
	ins structs_queries.RelationalOperator,
) (conditions.RelationalOperator, error) {
	builder := app.relationalOperatorBuilder.Create()
	if ins.And {
		builder.IsAnd()
	}

	if ins.Or {
		builder.IsOr()
	}

	return builder.Now()
}

func (app *resourceTokenQueryAdapter) integerOperatorToStruct(
	ins conditions.IntegerOperator,
) structs_queries.IntegerOperator {
	output := structs_queries.IntegerOperator{}
	if ins.IsSmallerThan() {
		output.SmallerThan = ins.IsSmallerThan()
	}

	if ins.IsBiggerThan() {
		output.BiggerThan = ins.IsBiggerThan()
	}

	if ins.IsEqual() {
		output.Equal = ins.IsEqual()
	}

	return output
}

func (app *resourceTokenQueryAdapter) structToIntegerOperator(
	ins structs_queries.IntegerOperator,
) (conditions.IntegerOperator, error) {
	builder := app.integerOperatorBuilder.Create()
	if ins.SmallerThan {
		builder.IsSmallerThan()
	}

	if ins.BiggerThan {
		builder.IsBiggerThan()
	}

	if ins.Equal {
		builder.IsEqual()
	}

	return builder.Now()
}
