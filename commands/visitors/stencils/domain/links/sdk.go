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
	WithWeight(weight uint) ElementBuilder
	WithContainer(container []string) ElementBuilder
	OnFailure(onFailure Element) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	Weight() uint
	Container() []string
	HasOnFailure() bool
	OnFailure() Element
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
