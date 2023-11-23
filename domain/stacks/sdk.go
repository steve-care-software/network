package stacks

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/layers"
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
	HasInstructions() bool
	Instructions() Instructions
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
	WithInstructions(instructions Instructions) FrameBuilder
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
	HasInstructions() bool
	Instructions() Instructions
}

// InstructionsBuilder represents an instructions builder
type InstructionsBuilder interface {
	Create() InstructionsBuilder
	WithList(list []Instruction) InstructionsBuilder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithSave(save layers.Layer) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsSave() bool
	Save() layers.Layer
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
