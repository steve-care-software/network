package stacks

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
)

// Builder represents a stack builder
type Builder interface {
	Create() Builder
	WithFrames(frames Frames) Builder
	Now() (Stack, error)
}

// Stack represents a stack
type Stack interface {
	Frames() Frames
	Body() Frames
	Last() Frame
}

// FramesBuilder represents the frames builder
type FramesBuilder interface {
	Create() FramesBuilder
	WithList(list []Frame) FramesBuilder
	Now() (Frames, error)
}

// Frames represents frames
type Frames interface {
	List() []Frame
}

// FrameBuilder represents the frame builder
type FrameBuilder interface {
	Create() FrameBuilder
	WithAssignments(assignments Assignments) FrameBuilder
	Now() (Frame, error)
}

// Frame represents a frame
type Frame interface {
	Fetch(name string) (Assignable, error)
	FetchBool(name string) (bool, error)
	FetchSignerPublicKeys(name string) ([]signers.PublicKey, error)
	FetchVote(name string) (signers.Vote, error)
	FetchSignature(name string) (signers.Signature, error)
	FetchHashList(name string) ([]hash.Hash, error)
	HasAssignments() bool
	Assignments() Assignments
}

// AssignmentsBuilder represents an assignments builder
type AssignmentsBuilder interface {
	Create() AssignmentsBuilder
	WithList(list []Assignment) AssignmentsBuilder
	Now() (Assignments, error)
}

// Assignments represents assignments
type Assignments interface {
	List() []Assignment
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithName(name string) AssignmentBuilder
	WithAssignable(assignable Assignable) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Name() string
	Assignable() Assignable
}

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithBool(boolValue bool) AssignableBuilder
	WithBytes(bytes []byte) AssignableBuilder
	WithSignerPublicKey(signerPublicKey signers.PublicKey) AssignableBuilder
	WithSignerPublicKeys(signerPubKeys []signers.PublicKey) AssignableBuilder
	WithHashList(hashList []hash.Hash) AssignableBuilder
	WithSignature(signature signers.Signature) AssignableBuilder
	WithVote(vote signers.Vote) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsBool() bool
	Bool() *bool
	IsBytes() bool
	Bytes() []byte
	IsSignerPublicKey() bool
	SignerPublicKey() signers.PublicKey
	IsSignerPublicKeys() bool
	SignerPublicKeys() []signers.PublicKey
	IsHashList() bool
	HashList() []hash.Hash
	IsSignature() bool
	Signature() signers.Signature
	IsVote() bool
	Vote() signers.Vote
}
