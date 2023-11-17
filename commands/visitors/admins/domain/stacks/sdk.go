package stacks

import (
	"steve.care/network/commands/visitors/admins/domain/accounts"
	"steve.care/network/libraries/hash"
)

const (
	// AuthorizedAccountDoNotContainIdentitiesError represents the authorized account do not contain identities error
	AuthorizedAccountDoNotContainIdentitiesError (uint) = iota

	// represents the account name already exists error
	AccountNameAlreadyExists
)

// Builder represents the stack builder
type Builder interface {
	Create() Builder
	WithMemory(memory Memory) Builder
	WithFrames(frames Frames) Builder
	Now() (Stack, error)
}

// Stack represents a stack
type Stack interface {
	Hash() hash.Hash
	Memory() Memory
	List() Frames
	Body() Frames
	Last() Frame
	ContainsError() bool
}

// MemoryBuilder represents a memory builder
type MemoryBuilder interface {
	Create() MemoryBuilder
	WithAuthorized(authorized accounts.Account) MemoryBuilder
	Now() (Memory, error)
}

// Memory represents memory
type Memory interface {
	Authorized() accounts.Account
}

// FramesBuilder repreents frames builder
type FramesBuilder interface {
	Create() FramesBuilder
	WithList(list []Frame) FramesBuilder
	Now() (Frames, error)
}

// Frames represents frames
type Frames interface {
	List() []Frame
}

// FrameFactory represents a frame factory
type FrameFactory interface {
	Create() Frame
}

// FrameBuilder represents a frame builder
type FrameBuilder interface {
	Create() FrameBuilder
	WithAssignments(assignments []Assignment) FrameBuilder
	WithInstructions(instructions []Instruction) FrameBuilder
	Now() (Frame, error)
}

// Frame represents frame
type Frame interface {
	Hash() hash.Hash
	Fetch(name string) (Assignable, error)
	FetchUint(name string) (*uint, error)
	ContainsError() bool
	HasAssignments() bool
	Assignments() Assignments
	HasInstructions() bool
	Instructions() Instructions
}

// InstructionsBuilder represents instructions builder
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
	WithCreateAccount(createAccount CreateAccount) InstructionBuilder
	IsDeleteAuthorized() InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents a frame instruction
type Instruction interface {
	IsDeleteAuthorized() bool
	IsCreateAccount() bool
	CreateAccount() CreateAccount
}

// CreateAccountBuilder represents a create account builder
type CreateAccountBuilder interface {
	Create() CreateAccountBuilder
	WithAccount(account accounts.Account) CreateAccountBuilder
	WithPassword(password []byte) CreateAccountBuilder
	Now() (CreateAccount, error)
}

// CreateAccount represents a create account
type CreateAccount interface {
	Account() accounts.Account
	Password() []byte
}

// AssignmentsBuilder represents assignments builder
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
	Hash() hash.Hash
	Name() string
	Assignable() Assignable
}

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithBool(boolValue bool) AssignableBuilder
	WithIdentities(identities accounts.Identities) AssignableBuilder
	WithError(errorCode uint) AssignableBuilder
	Now() (Assignable, error)
}

type Assignable interface {
	Hash() hash.Hash
	IsBool() bool
	Bool() *bool
	IsIdentities() bool
	Identities() accounts.Identities
	IsError() bool
	Error() uint
}
