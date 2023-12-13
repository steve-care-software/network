package links

import (
	"errors"

	"steve.care/network/domain/receipts/commands/links"
)

type builder struct {
	link              links.Link
	element           links.Element
	condition         links.Condition
	conditionValue    links.ConditionValue
	conditionResource links.ConditionResource
	origin            links.Origin
	originValue       links.OriginValue
	originResource    links.OriginResource
	operator          links.Operator
}

func createBuilder() Builder {
	out := builder{
		link:              nil,
		element:           nil,
		condition:         nil,
		conditionValue:    nil,
		conditionResource: nil,
		origin:            nil,
		originValue:       nil,
		originResource:    nil,
		operator:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithLink adds a link to the builder
func (app *builder) WithLink(link links.Link) Builder {
	app.link = link
	return app
}

// WithElement adds a link to the builder
func (app *builder) WithElement(element links.Element) Builder {
	app.element = element
	return app
}

// WithCondition adds a condition to the builder
func (app *builder) WithCondition(condition links.Condition) Builder {
	app.condition = condition
	return app
}

// WithConditionValue adds a conditionValue to the builder
func (app *builder) WithConditionValue(conditionValue links.ConditionValue) Builder {
	app.conditionValue = conditionValue
	return app
}

// WithConditionResource adds a conditionResource to the builder
func (app *builder) WithConditionResource(conditionResource links.ConditionResource) Builder {
	app.conditionResource = conditionResource
	return app
}

// WithOrigin adds an origin to the builder
func (app *builder) WithOrigin(origin links.Origin) Builder {
	app.origin = origin
	return app
}

// WithOriginValue adds an originValue to the builder
func (app *builder) WithOriginValue(originValue links.OriginValue) Builder {
	app.originValue = originValue
	return app
}

// WithOriginResource adds an originResource to the builder
func (app *builder) WithOriginResource(originResource links.OriginResource) Builder {
	app.originResource = originResource
	return app
}

// WithOperator adds an operator to the builder
func (app *builder) WithOperator(operator links.Operator) Builder {
	app.operator = operator
	return app
}

// Now builds a new Link instance
func (app *builder) Now() (Link, error) {
	if app.link != nil {
		return createLinkWithLink(app.link), nil
	}

	if app.element != nil {
		return createLinkWithElement(app.element), nil
	}

	if app.condition != nil {
		return createLinkWithCondition(app.condition), nil
	}

	if app.conditionValue != nil {
		return createLinkWithConditionValue(app.conditionValue), nil
	}

	if app.conditionResource != nil {
		return createLinkWithConditionResource(app.conditionResource), nil
	}

	if app.origin != nil {
		return createLinkWithOrigin(app.origin), nil
	}

	if app.originValue != nil {
		return createLinkWithOriginValue(app.originValue), nil
	}

	if app.originResource != nil {
		return createLinkWithOriginResource(app.originResource), nil
	}

	if app.operator != nil {
		return createLinkWithOperator(app.operator), nil
	}

	return nil, errors.New("the Link resource is invalid")
}
