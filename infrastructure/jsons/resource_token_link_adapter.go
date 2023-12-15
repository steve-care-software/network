package jsons

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/links"
	structs_links "steve.care/network/infrastructure/jsons/resources/tokens/links"
)

type resourceTokenLinkAdapter struct {
	hashAdapter           hash.Adapter
	originBuilder         links.OriginBuilder
	originValueBuilder    links.OriginValueBuilder
	originResourceBuilder links.OriginResourceBuilder
	operatorBuilder       links.OperatorBuilder
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

func (app *resourceTokenLinkAdapter) originToStruct(
	ins links.Origin,
) structs_links.Origin {
	resource := app.originResourceToStruct(ins.Resource())
	operator := app.operatorToStruct(ins.Operator())
	next := app.originValueToStruct(ins.Next())
	return structs_links.Origin{
		Resource: resource,
		Operator: operator,
		Next:     next,
	}
}

func (app *resourceTokenLinkAdapter) structToOrigin(
	ins structs_links.Origin,
) (links.Origin, error) {
	resource, err := app.structToOriginResource(ins.Resource)
	if err != nil {
		return nil, err
	}

	operator, err := app.structToOperator(ins.Operator)
	if err != nil {
		return nil, err
	}

	next, err := app.structToOriginValue(ins.Next)
	if err != nil {
		return nil, err
	}

	return app.originBuilder.Create().
		WithResource(resource).
		WithOperator(operator).
		WithNext(next).
		Now()
}

func (app *resourceTokenLinkAdapter) originValueToStruct(
	ins links.OriginValue,
) structs_links.OriginValue {
	output := structs_links.OriginValue{}
	if ins.IsResource() {
		resource := app.originResourceToStruct(ins.Resource())
		output.Resource = &resource
	}

	if ins.IsOrigin() {
		origin := app.originToStruct(ins.Origin())
		output.Origin = &origin
	}

	return output
}

func (app *resourceTokenLinkAdapter) structToOriginValue(
	ins structs_links.OriginValue,
) (links.OriginValue, error) {
	builder := app.originValueBuilder.Create()
	if ins.Resource != nil {
		resource, err := app.structToOriginResource(*ins.Resource)
		if err != nil {
			return nil, err
		}

		builder.WithResource(resource)
	}

	if ins.Origin != nil {
		origin, err := app.structToOrigin(*ins.Origin)
		if err != nil {
			return nil, err
		}

		builder.WithOrigin(origin)
	}

	return builder.Now()
}

func (app *resourceTokenLinkAdapter) originResourceToStruct(
	ins links.OriginResource,
) structs_links.OriginResource {
	return structs_links.OriginResource{
		Layer:       ins.Hash().String(),
		IsMandatory: ins.IsMandatory(),
	}
}

func (app *resourceTokenLinkAdapter) structToOriginResource(
	ins structs_links.OriginResource,
) (links.OriginResource, error) {
	pLayerHash, err := app.hashAdapter.FromString(ins.Layer)
	if err != nil {
		return nil, err
	}

	builder := app.originResourceBuilder.Create().
		WithLayer(*pLayerHash)

	if ins.IsMandatory {
		builder.IsMandatory()
	}

	return builder.Now()
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
		builder.IsAnd()
	}

	if ins.Or {
		builder.IsOr()
	}

	if ins.Xor {
		builder.IsXor()
	}

	return builder.Now()
}
