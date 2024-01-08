package links

import (
	"reflect"
	"testing"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries/layers/links"
)

func TestLink_withLink_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))

	link := links.NewLinkForTests(
		links.NewOriginForTests(
			links.NewOriginResourceForTests(*pFirstLayer),
			links.NewOperatorWithAndForTests(),
			links.NewOriginValueWithResourceForTests(
				links.NewOriginResourceForTests(*pSecondLayer),
			),
		),
		links.NewElementsForTests([]links.Element{
			links.NewElementForTests(*pLayer),
		}),
	)

	ins := NewLinkWithLinkForTests(link)

	if !ins.IsLink() {
		t.Errorf("the link was expected to contain a link")
		return
	}

	if ins.IsElement() {
		t.Errorf("the link was expected to NOT contain an element")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the link was expected to NOT contain a condition")
		return
	}

	if ins.IsConditionValue() {
		t.Errorf("the link was expected to NOT contain a conditionValue")
		return
	}

	if ins.IsConditionResource() {
		t.Errorf("the link was expected to NOT contain a conditionResource")
		return
	}

	if ins.IsOrigin() {
		t.Errorf("the link was expected to NOT contain an origin")
		return
	}

	if ins.IsOriginValue() {
		t.Errorf("the link was expected to NOT contain an originValue")
		return
	}

	if ins.IsOriginResource() {
		t.Errorf("the link was expected to NOT contain an originResource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the link was expected to NOT contain an operator")
		return
	}

	retLink := ins.Link()
	if !reflect.DeepEqual(link, retLink) {
		t.Errorf("the returned link is invalid")
		return
	}
}

func TestLink_withElement_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	element := links.NewElementForTests(*pLayer)
	ins := NewLinkWithElementForTests(element)

	if ins.IsLink() {
		t.Errorf("the link was expected to NOT contain a link")
		return
	}

	if !ins.IsElement() {
		t.Errorf("the link was expected to contain an element")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the link was expected to NOT contain a condition")
		return
	}

	if ins.IsConditionValue() {
		t.Errorf("the link was expected to NOT contain a conditionValue")
		return
	}

	if ins.IsConditionResource() {
		t.Errorf("the link was expected to NOT contain a conditionResource")
		return
	}

	if ins.IsOrigin() {
		t.Errorf("the link was expected to NOT contain an origin")
		return
	}

	if ins.IsOriginValue() {
		t.Errorf("the link was expected to NOT contain an originValue")
		return
	}

	if ins.IsOriginResource() {
		t.Errorf("the link was expected to NOT contain an originResource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the link was expected to NOT contain an operator")
		return
	}

	retElement := ins.Element()
	if !reflect.DeepEqual(element, retElement) {
		t.Errorf("the returned element is invalid")
		return
	}
}

func TestLink_withCondition_Success(t *testing.T) {
	condition := links.NewConditionForTests(
		links.NewConditionResourceForTests(23),
		links.NewOperatorWithAndForTests(),
		links.NewConditionValueWithResourceForTests(
			links.NewConditionResourceForTests(44),
		),
	)

	ins := NewLinkWithConditionForTests(condition)

	if ins.IsLink() {
		t.Errorf("the link was expected to NOT contain a link")
		return
	}

	if ins.IsElement() {
		t.Errorf("the link was expected to NOT contain an element")
		return
	}

	if !ins.IsCondition() {
		t.Errorf("the link was expected to contain a condition")
		return
	}

	if ins.IsConditionValue() {
		t.Errorf("the link was expected to NOT contain a conditionValue")
		return
	}

	if ins.IsConditionResource() {
		t.Errorf("the link was expected to NOT contain a conditionResource")
		return
	}

	if ins.IsOrigin() {
		t.Errorf("the link was expected to NOT contain an origin")
		return
	}

	if ins.IsOriginValue() {
		t.Errorf("the link was expected to NOT contain an originValue")
		return
	}

	if ins.IsOriginResource() {
		t.Errorf("the link was expected to NOT contain an originResource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the link was expected to NOT contain an operator")
		return
	}

	retCondition := ins.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the returned condition is invalid")
		return
	}
}

func TestLink_withConditionValue_Success(t *testing.T) {
	conditionValue := links.NewConditionValueWithResourceForTests(
		links.NewConditionResourceForTests(45),
	)

	ins := NewLinkWithConditionValueForTests(conditionValue)

	if ins.IsLink() {
		t.Errorf("the link was expected to NOT contain a link")
		return
	}

	if ins.IsElement() {
		t.Errorf("the link was expected to NOT contain an element")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the link was expected to NOT contain a condition")
		return
	}

	if !ins.IsConditionValue() {
		t.Errorf("the link was expected to contain a conditionValue")
		return
	}

	if ins.IsConditionResource() {
		t.Errorf("the link was expected to NOT contain a conditionResource")
		return
	}

	if ins.IsOrigin() {
		t.Errorf("the link was expected to NOT contain an origin")
		return
	}

	if ins.IsOriginValue() {
		t.Errorf("the link was expected to NOT contain an originValue")
		return
	}

	if ins.IsOriginResource() {
		t.Errorf("the link was expected to NOT contain an originResource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the link was expected to NOT contain an operator")
		return
	}

	retConditionValue := ins.ConditionValue()
	if !reflect.DeepEqual(conditionValue, retConditionValue) {
		t.Errorf("the returned conditionValue is invalid")
		return
	}
}

func TestLink_withConditionResource_Success(t *testing.T) {
	conditionResource := links.NewConditionResourceForTests(45)
	ins := NewLinkWithConditionResourceForTests(conditionResource)

	if ins.IsLink() {
		t.Errorf("the link was expected to NOT contain a link")
		return
	}

	if ins.IsElement() {
		t.Errorf("the link was expected to NOT contain an element")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the link was expected to NOT contain a condition")
		return
	}

	if ins.IsConditionValue() {
		t.Errorf("the link was expected to NOT contain a conditionValue")
		return
	}

	if !ins.IsConditionResource() {
		t.Errorf("the link was expected to contain a conditionResource")
		return
	}

	if ins.IsOrigin() {
		t.Errorf("the link was expected to NOT contain an origin")
		return
	}

	if ins.IsOriginValue() {
		t.Errorf("the link was expected to NOT contain an originValue")
		return
	}

	if ins.IsOriginResource() {
		t.Errorf("the link was expected to NOT contain an originResource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the link was expected to NOT contain an operator")
		return
	}

	retConditionResource := ins.ConditionResource()
	if !reflect.DeepEqual(conditionResource, retConditionResource) {
		t.Errorf("the returned conditionResource is invalid")
		return
	}
}

func TestLink_withOrigin_Success(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	origin := links.NewOriginForTests(
		links.NewOriginResourceForTests(*pFirstLayer),
		links.NewOperatorWithAndForTests(),
		links.NewOriginValueWithResourceForTests(
			links.NewOriginResourceForTests(*pSecondLayer),
		),
	)

	ins := NewLinkWithOriginForTests(origin)

	if ins.IsLink() {
		t.Errorf("the link was expected to NOT contain a link")
		return
	}

	if ins.IsElement() {
		t.Errorf("the link was expected to NOT contain an element")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the link was expected to NOT contain a condition")
		return
	}

	if ins.IsConditionValue() {
		t.Errorf("the link was expected to NOT contain a conditionValue")
		return
	}

	if ins.IsConditionResource() {
		t.Errorf("the link was expected to NOT contain a conditionResource")
		return
	}

	if !ins.IsOrigin() {
		t.Errorf("the link was expected to contain an origin")
		return
	}

	if ins.IsOriginValue() {
		t.Errorf("the link was expected to NOT contain an originValue")
		return
	}

	if ins.IsOriginResource() {
		t.Errorf("the link was expected to NOT contain an originResource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the link was expected to NOT contain an operator")
		return
	}

	retOrigin := ins.Origin()
	if !reflect.DeepEqual(origin, retOrigin) {
		t.Errorf("the returned origin is invalid")
		return
	}
}

func TestLink_withOriginValue_Success(t *testing.T) {
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	originValue := links.NewOriginValueWithResourceForTests(
		links.NewOriginResourceForTests(*pSecondLayer),
	)

	ins := NewLinkWithOriginValueForTests(originValue)

	if ins.IsLink() {
		t.Errorf("the link was expected to NOT contain a link")
		return
	}

	if ins.IsElement() {
		t.Errorf("the link was expected to NOT contain an element")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the link was expected to NOT contain a condition")
		return
	}

	if ins.IsConditionValue() {
		t.Errorf("the link was expected to NOT contain a conditionValue")
		return
	}

	if ins.IsConditionResource() {
		t.Errorf("the link was expected to NOT contain a conditionResource")
		return
	}

	if ins.IsOrigin() {
		t.Errorf("the link was expected to NOT contain an origin")
		return
	}

	if !ins.IsOriginValue() {
		t.Errorf("the link was expected to contain an originValue")
		return
	}

	if ins.IsOriginResource() {
		t.Errorf("the link was expected to NOT contain an originResource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the link was expected to NOT contain an operator")
		return
	}

	retOriginValue := ins.OriginValue()
	if !reflect.DeepEqual(originValue, retOriginValue) {
		t.Errorf("the returned originValue is invalid")
		return
	}
}

func TestLink_withOriginResource_Success(t *testing.T) {
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	originResource := links.NewOriginResourceForTests(*pSecondLayer)

	ins := NewLinkWithOriginResourceForTests(originResource)

	if ins.IsLink() {
		t.Errorf("the link was expected to NOT contain a link")
		return
	}

	if ins.IsElement() {
		t.Errorf("the link was expected to NOT contain an element")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the link was expected to NOT contain a condition")
		return
	}

	if ins.IsConditionValue() {
		t.Errorf("the link was expected to NOT contain a conditionValue")
		return
	}

	if ins.IsConditionResource() {
		t.Errorf("the link was expected to NOT contain a conditionResource")
		return
	}

	if ins.IsOrigin() {
		t.Errorf("the link was expected to NOT contain an origin")
		return
	}

	if ins.IsOriginValue() {
		t.Errorf("the link was expected to NOT contain an originValue")
		return
	}

	if !ins.IsOriginResource() {
		t.Errorf("the link was expected to contain an originResource")
		return
	}

	if ins.IsOperator() {
		t.Errorf("the link was expected to NOT contain an operator")
		return
	}

	retOriginResource := ins.OriginResource()
	if !reflect.DeepEqual(originResource, retOriginResource) {
		t.Errorf("the returned originResource is invalid")
		return
	}
}

func TestLink_withOperator_Success(t *testing.T) {
	operator := links.NewOperatorWithAndForTests()
	ins := NewLinkWithOperatorForTests(operator)

	if ins.IsLink() {
		t.Errorf("the link was expected to NOT contain a link")
		return
	}

	if ins.IsElement() {
		t.Errorf("the link was expected to NOT contain an element")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the link was expected to NOT contain a condition")
		return
	}

	if ins.IsConditionValue() {
		t.Errorf("the link was expected to NOT contain a conditionValue")
		return
	}

	if ins.IsConditionResource() {
		t.Errorf("the link was expected to NOT contain a conditionResource")
		return
	}

	if ins.IsOrigin() {
		t.Errorf("the link was expected to NOT contain an origin")
		return
	}

	if ins.IsOriginValue() {
		t.Errorf("the link was expected to NOT contain an originValue")
		return
	}

	if ins.IsOriginResource() {
		t.Errorf("the link was expected to NOT contain an originResource")
		return
	}

	if !ins.IsOperator() {
		t.Errorf("the link was expected to contain an operator")
		return
	}

	retOperator := ins.Operator()
	if !reflect.DeepEqual(operator, retOperator) {
		t.Errorf("the returned operator is invalid")
		return
	}
}

func TestLink_withoutParam_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
