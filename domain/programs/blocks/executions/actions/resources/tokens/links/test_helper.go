package links

import "steve.care/network/domain/programs/logics/libraries/layers/links"

// NewLinkWithLinkForTests creates a new link with link for tests
func NewLinkWithLinkForTests(input links.Link) Link {
	ins, err := NewBuilder().Create().WithLink(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithElementForTests creates a new link with element for tests
func NewLinkWithElementForTests(input links.Element) Link {
	ins, err := NewBuilder().Create().WithElement(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithConditionForTests creates a new link with condition for tests
func NewLinkWithConditionForTests(input links.Condition) Link {
	ins, err := NewBuilder().Create().WithCondition(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithConditionValueForTests creates a new link with conditionValue for tests
func NewLinkWithConditionValueForTests(input links.ConditionValue) Link {
	ins, err := NewBuilder().Create().WithConditionValue(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithConditionResourceForTests creates a new link with conditionResource for tests
func NewLinkWithConditionResourceForTests(input links.ConditionResource) Link {
	ins, err := NewBuilder().Create().WithConditionResource(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithOriginForTests creates a new link with origin for tests
func NewLinkWithOriginForTests(input links.Origin) Link {
	ins, err := NewBuilder().Create().WithOrigin(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithOriginValueForTests creates a new link with originValue for tests
func NewLinkWithOriginValueForTests(input links.OriginValue) Link {
	ins, err := NewBuilder().Create().WithOriginValue(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithOriginResourceForTests creates a new link with originResource for tests
func NewLinkWithOriginResourceForTests(input links.OriginResource) Link {
	ins, err := NewBuilder().Create().WithOriginResource(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithOperatorForTests creates a new link with operator for tests
func NewLinkWithOperatorForTests(input links.Operator) Link {
	ins, err := NewBuilder().Create().WithOperator(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
