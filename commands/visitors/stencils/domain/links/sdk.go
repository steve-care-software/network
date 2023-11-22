package links

import (
	"steve.care/network/libraries/hash"
)

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
	WithValue(value ConditionValue) ConditionBuilder
	WithOperator(operator Operator) ConditionBuilder
	WithNext(next ConditionNext) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Value() ConditionValue
	Operator() Operator
	Next() ConditionNext
}

// ConditionNextBuilder represents a condition next builder
type ConditionNextBuilder interface {
	Create() ConditionNextBuilder
	WithValue(value ConditionValue) ConditionNextBuilder
	WithCondition(condition Condition) ConditionNextBuilder
	Now() (ConditionNext, error)
}

// ConditionNext represents a condition next
type ConditionNext interface {
	Hash() hash.Hash
	IsValue() bool
	Value() ConditionValue
	IsCondition() bool
	Condition() Condition
}

// ConditionValueBuilder represents a condition value builder
type ConditionValueBuilder interface {
	Create() ConditionValueBuilder
	WithCode(code uint) ConditionValueBuilder
	IsRaisedInLayer() ConditionValueBuilder
	Now() (ConditionValue, error)
}

// ConditionValue represents a condition value
type ConditionValue interface {
	Hash() hash.Hash
	Code() uint
	IsRaisedInLayer() bool
}

// OriginBuilder represents the origin builder
type OriginBuilder interface {
	Create() OriginBuilder
	WithLayer(layer Layer) OriginBuilder
	WithOperator(operator Operator) OriginBuilder
	WithNext(next Resource) OriginBuilder
	Now() (Origin, error)
}

// Origin represents an origin
type Origin interface {
	Hash() hash.Hash
	Layer() Layer
	Operator() Operator
	Next() Resource
}

// ResourceBuilder represents the resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithLayer(layer Layer) ResourceBuilder
	WithOrigin(origin Origin) ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	IsLayer() bool
	Layer() Layer
	IsOrigin() bool
	Origin() Origin
}

// LayerBuilder represents the layer builder
type LayerBuilder interface {
	Create() LayerBuilder
	WithContainer(container []string) LayerBuilder
	IsMandatory() LayerBuilder
	Now() (Layer, error)
}

// Layer represents a layer
type Layer interface {
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
