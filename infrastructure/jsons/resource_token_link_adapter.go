package jsons

import (
	"steve.care/network/domain/receipts/commands/links"
	structs_links "steve.care/network/infrastructure/jsons/resources/tokens/links"
)

type resourceTokenLinkAdapter struct {
	operatorBuilder links.OperatorBuilder
}

// LinkToStruct converts a link to struct
func (app *resourceTokenLinkAdapter) LinkToStruct(
	ins links.Link,
) structs_links.Link {
	return structs_links.Link{}
}

// StructToLink converts a struct to link
func (app *resourceTokenLinkAdapter) StructToLink(
	ins structs_links.Link,
) (links.Link, error) {
	return nil, nil
}

func (app *resourceTokenLinkAdapter) operatorToStruct(
	ins links.Operator,
) structs_links.Operator {
	output := structs_links.Operator{}
	if ins.IsAnd() {
		output.And = ins.IsAnd()
	}

	if ins.IsOr() {
		output.Or = ins.IsOr()
	}

	if ins.IsXor() {
		output.Xor = ins.IsXor()
	}

	return output
}

func (app *resourceTokenLinkAdapter) structToOperator(
	ins structs_links.Operator,
) (links.Operator, error) {
	builder := app.operatorBuilder.Create()
	if ins.And {

	}

	if ins.Or {

	}

	if ins.Xor {

	}

	return builder.Now()
}
