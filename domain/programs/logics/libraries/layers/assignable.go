package layers

import "steve.care/network/domain/hash"

type assignable struct {
	hash     hash.Hash
	bytes    Bytes
	identity Identity
	engine   Engine
}

func createAssignableWithBytes(
	hash hash.Hash,
	bytes Bytes,
) Assignable {
	return createAssignableInternally(hash, bytes, nil, nil)
}

func createAssignableWithIdentity(
	hash hash.Hash,
	identity Identity,
) Assignable {
	return createAssignableInternally(hash, nil, identity, nil)
}

func createAssignableWithEngine(
	hash hash.Hash,
	engine Engine,
) Assignable {
	return createAssignableInternally(hash, nil, nil, engine)
}

func createAssignableInternally(
	hash hash.Hash,
	bytes Bytes,
	identity Identity,
	engine Engine,
) Assignable {
	out := assignable{
		hash:     hash,
		bytes:    bytes,
		identity: identity,
		engine:   engine,
	}

	return &out
}

// Hash returns the hash
func (obj *assignable) Hash() hash.Hash {
	return obj.hash
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *assignable) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *assignable) Bytes() Bytes {
	return obj.bytes
}

// IsIdentity returns true if there is identity, false otherwise
func (obj *assignable) IsIdentity() bool {
	return obj.identity != nil
}

// Identity returns the identity, if any
func (obj *assignable) Identity() Identity {
	return obj.identity
}

// IsEngine returns true if there is engine, false otherwise
func (obj *assignable) IsEngine() bool {
	return obj.engine != nil
}

// Engine returns the engine, if any
func (obj *assignable) Engine() Engine {
	return obj.engine
}
