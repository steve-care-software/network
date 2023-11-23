package links

import (
	"steve.care/network/libraries/hash"
)

// NewElementsBuilder creates a new elements builder
func NewElementsBuilder() ElementsBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementsBuilder(
		hashAdapter,
	)
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementBuilder(
		hashAdapter,
	)
}

// NewConditionBuilder creates a new condition builder
func NewConditionBuilder() ConditionBuilder {
	hashAdapter := hash.NewAdapter()
	return createConditionBuilder(
		hashAdapter,
	)
}

// NewConditionValueBuilder creates a new condition value builder
func NewConditionValueBuilder() ConditionValueBuilder {
	hashAdapter := hash.NewAdapter()
	return createConditionValueBuilder(
		hashAdapter,
	)
}

// NewConditionResourceBuilder creates a new condition resource builder
func NewConditionResourceBuilder() ConditionResourceBuilder {
	hashAdapter := hash.NewAdapter()
	return createConditionResourceBuilder(
		hashAdapter,
	)
}

// NewOriginBuilder creates a new origin builder
func NewOriginBuilder() OriginBuilder {
	hashAdapter := hash.NewAdapter()
	return createOriginBuilder(
		hashAdapter,
	)
}

// NewOriginValueBuilder creates a new origin value builder
func NewOriginValueBuilder() OriginValueBuilder {
	hashAdapter := hash.NewAdapter()
	return createOriginValueBuilder(
		hashAdapter,
	)
}

// NewOriginResourceBuilder creates a new origin resource builder
func NewOriginResourceBuilder() OriginResourceBuilder {
	hashAdapter := hash.NewAdapter()
	return createOriginResourceBuilder(
		hashAdapter,
	)
}

// NewOperatorBuilder creates a new operator builder
func NewOperatorBuilder() OperatorBuilder {
	hashAdapter := hash.NewAdapter()
	return createOperatorBuilder(
		hashAdapter,
	)
}

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithOrigin(origin Origin) Builder
	WithElements(elements Elements) Builder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	Hash() hash.Hash
	Origin() Origin
	Elements() Elements
}

// ElementsBuilder represents elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithList(list []Element) ElementsBuilder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	Hash() hash.Hash
	List() []Element
}

// ElementBuilder represents the element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithContainer(container []string) ElementBuilder
	WithCondition(condition Condition) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	Container() []string
	HasCondition() bool
	Condition() Condition
}

// ConditionBuilder represents condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithResource(resource ConditionResource) ConditionBuilder
	WithOperator(operator Operator) ConditionBuilder
	WithNext(next ConditionValue) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Resource() ConditionResource
	Operator() Operator
	Next() ConditionValue
}

// ConditionValueBuilder represents a condition value builder
type ConditionValueBuilder interface {
	Create() ConditionValueBuilder
	WithResource(resource ConditionResource) ConditionValueBuilder
	WithCondition(condition Condition) ConditionValueBuilder
	Now() (ConditionValue, error)
}

// ConditionValue represents a condition value
type ConditionValue interface {
	Hash() hash.Hash
	IsResource() bool
	Resource() ConditionResource
	IsCondition() bool
	Condition() Condition
}

// ConditionResourceBuilder represents a condition resource builder
type ConditionResourceBuilder interface {
	Create() ConditionResourceBuilder
	WithCode(code uint) ConditionResourceBuilder
	IsRaisedInLayer() ConditionResourceBuilder
	Now() (ConditionResource, error)
}

// ConditionResource represents a condition resource
type ConditionResource interface {
	Hash() hash.Hash
	Code() uint
	IsRaisedInLayer() bool
}

// OriginBuilder represents the origin builder
type OriginBuilder interface {
	Create() OriginBuilder
	WithResource(resource OriginResource) OriginBuilder
	WithOperator(operator Operator) OriginBuilder
	WithNext(next OriginValue) OriginBuilder
	Now() (Origin, error)
}

// Origin represents an origin
type Origin interface {
	Hash() hash.Hash
	Resource() OriginResource
	Operator() Operator
	Next() OriginValue
}

// OriginValueBuilder represents the originValue builder
type OriginValueBuilder interface {
	Create() OriginValueBuilder
	WithResource(resource OriginResource) OriginValueBuilder
	WithOrigin(origin Origin) OriginValueBuilder
	Now() (OriginValue, error)
}

// OriginValue represents an origin value
type OriginValue interface {
	Hash() hash.Hash
	IsResource() bool
	Resource() OriginResource
	IsOrigin() bool
	Origin() Origin
}

// OriginResourceBuilder represents the origin resource builder
type OriginResourceBuilder interface {
	Create() OriginResourceBuilder
	WithContainer(container []string) OriginResourceBuilder
	IsMandatory() OriginResourceBuilder
	Now() (OriginResource, error)
}

// OriginResource represents an origin resource
type OriginResource interface {
	Hash() hash.Hash
	Container() []string
	IsMandatory() bool
}

// OperatorBuilder represents the operator builder
type OperatorBuilder interface {
	Create() OperatorBuilder
	IsAnd() OperatorBuilder
	IsOr() OperatorBuilder
	IsXor() OperatorBuilder
	Now() (Operator, error)
}

// Operator represents the operator
type Operator interface {
	Hash() hash.Hash
	IsAnd() bool
	IsOr() bool
	IsXor() bool
}

// Repository represents the link repository
type Repository interface {
	Retrieve(executedLayers [][]string) (Link, error)
}

// Service represents the link service
type Service interface {
	Insert(link Link) error
	Update(origin hash.Hash, updated Link) error
	Delete(hash hash.Hash) error
}
