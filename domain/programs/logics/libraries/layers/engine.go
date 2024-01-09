package layers

import "steve.care/network/domain/hash"

type engine struct {
	hash      hash.Hash
	execution Execution
	resource  AssignableResource
}

func createEngineWithExecution(
	hash hash.Hash,
	execution Execution,
) Engine {
	return createEngineInternally(hash, execution, nil)
}

func createEngineWithResource(
	hash hash.Hash,
	resource AssignableResource,
) Engine {
	return createEngineInternally(hash, nil, resource)
}

func createEngineInternally(
	hash hash.Hash,
	execution Execution,
	resource AssignableResource,
) Engine {
	out := engine{
		hash:      hash,
		execution: execution,
		resource:  resource,
	}

	return &out
}

// Hash returns the hash
func (obj *engine) Hash() hash.Hash {
	return obj.hash
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *engine) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *engine) Execution() Execution {
	return obj.execution
}

// IsResource returns true if there is a resource, false otherwise
func (obj *engine) IsResource() bool {
	return obj.resource != nil
}

// Resource returns the resource, if any
func (obj *engine) Resource() AssignableResource {
	return obj.resource
}
