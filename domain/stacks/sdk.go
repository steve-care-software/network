package stacks

import (
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	framesBuilder := NewFramesBuilder()
	return createBuilder(
		framesBuilder,
	)
}

// NewFramesBuilder creates a new frames builder
func NewFramesBuilder() FramesBuilder {
	return createFramesBuilder()
}

// NewFrameBuilder creates a new frame builder
func NewFrameBuilder() FrameBuilder {
	return createFrameBuilder()
}

// NewAssignmentsBuilder creates a new assignments builder
func NewAssignmentsBuilder() AssignmentsBuilder {
	return createAssignmentsBuilder()
}

// NewAssignmentBuilder creates a new assignment builder
func NewAssignmentBuilder() AssignmentBuilder {
	return createAssignmentBuilder()
}

// NewAssignableBuilder creates a new assignable builder
func NewAssignableBuilder() AssignableBuilder {
	return createAssignableBuilder()
}

// Factory represents the stack factory
type Factory interface {
	Create() Stack
}

// Builder represents a stack builder
type Builder interface {
	Create() Builder
	WithFrames(frames Frames) Builder
	Now() (Stack, error)
}

// Stack represents a stack
type Stack interface {
	Frames() Frames
	Head() Frame
	HasBody() bool
	Body() Frames
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
	FetchHash(name string) (hash.Hash, error)
	FetchResource(name string) (resources.Resource, error)
	FetchBytes(name string) ([]byte, error)
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
	Fetch(name string) (Assignable, error)
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
	WithEncryptorPublicKey(encryptorPublicKey encryptors.PublicKey) AssignableBuilder
	WithSignerPublicKey(signerPublicKey signers.PublicKey) AssignableBuilder
	WithSignerPublicKeys(signerPubKeys []signers.PublicKey) AssignableBuilder
	WithSignature(signature signers.Signature) AssignableBuilder
	WithVote(vote signers.Vote) AssignableBuilder
	WithHashList(hashList []hash.Hash) AssignableBuilder
	WithHash(hash hash.Hash) AssignableBuilder
	WithResource(resource resources.Resource) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsBool() bool
	Bool() *bool
	IsBytes() bool
	Bytes() []byte
	IsEncryptorPublicKey() bool
	EncryptorPublicKey() encryptors.PublicKey
	IsSignerPublicKey() bool
	SignerPublicKey() signers.PublicKey
	IsSignerPublicKeys() bool
	SignerPublicKeys() []signers.PublicKey
	IsSignature() bool
	Signature() signers.Signature
	IsVote() bool
	Vote() signers.Vote
	IsHashList() bool
	HashList() []hash.Hash
	IsHash() bool
	Hash() hash.Hash
	IsResource() bool
	Resource() resources.Resource
}
