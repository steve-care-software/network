package stacks

import (
	admin_accounts "steve.care/network/commands/visitors/admins/domain/accounts"
	"steve.care/network/commands/visitors/admins/identities/dashboards/stencils/layers/domain/layers"
	identity_accounts "steve.care/network/commands/visitors/admins/identities/domain/accounts"
)

const (
	// LayerAlreadyExistsError represents the layer already exists error
	LayerAlreadyExistsError (int) = iota

	// LayerDoesNotExistsError represents the layer does not exists error
	LayerDoesNotExistsError
)

// Builder represents a stack builder
type Builder interface {
	Create() Builder
	WithMemory(memory Memory) Builder
	WithFrames(frames Frames) Builder
	Now() (Stack, error)
}

// Stack represents a stack
type Stack interface {
	Memory() Memory
	Frames() Frames
	Body() Frames
	Last() Frame
	HasInstructions() bool
	Instructions() Instructions
}

// MemoryBuilder represents a memory builder
type MemoryBuilder interface {
	Create() MemoryBuilder
	WithAuthorized(authorized admin_accounts.Account) MemoryBuilder
	WithAuthenticated(authenticated identity_accounts.Account) MemoryBuilder
	Now() (Memory, error)
}

// Memory represents a memory
type Memory interface {
	Authorized() admin_accounts.Account
	Authenticated() identity_accounts.Account
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
	HasInstructions() bool
	Instructions() Instructions
	HasAssignments() bool
	Assignments() Assignments
}

// InstructionsBuilder represents an instructions builder
type InstructionsBuilder interface {
	Create() InstructionsBuilder
	WithInstructions(instructions []Instruction) InstructionsBuilder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithSave(save Resource) InstructionBuilder
	WithMove(move Move) InstructionBuilder
	WithDelete(delete []string) InstructionBuilder
	WithErrorCode(errorCode int) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsSave() bool
	Save() Resource
	IsMove() bool
	Move() Move
	IsDelete() bool
	Delete() []string
	IsErrorCode() bool
	ErrorCode() *int
}

// ResourceBuilder represents a resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithPath(path []string) ResourceBuilder
	WithLayer(layer layers.Layer) ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Path() []string
	Layer() layers.Layer
}

// MoveBuilder represents a move builder
type MoveBuilder interface {
	Create() MoveBuilder
	From(from []string) MoveBuilder
	To(to []string) MoveBuilder
	Now() (Move, error)
}

// Move represents a move layer
type Move interface {
	From() []string
	To() []string
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

// AssignableBuilder represents an assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithBool(boolValue bool) AssignableBuilder
	WithStringList(strList []string) AssignableBuilder
	WithLayer(layer layers.Layer) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsBool() bool
	Bool() *bool
	IsStringList() bool
	StringList() []string
	IsLayer() bool
	Layer() layers.Layer
}
