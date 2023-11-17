package stacks

import (
	"steve.care/network/commands/visitors/admins/domain/accounts"
	admin_accounts "steve.care/network/commands/visitors/admins/domain/accounts"
	admin_stacks "steve.care/network/commands/visitors/admins/domain/stacks"
	identity_accounts "steve.care/network/commands/visitors/admins/identities/domain/accounts"
	stencil_stacks "steve.care/network/commands/visitors/stencils/domain/stacks"
	"steve.care/network/libraries/credentials"
	"steve.care/network/libraries/hash"
)

const (
	// CouldNotAuthorizeError represents a could not authorize error
	CouldNotAuthorizeError (uint) = iota

	// AccountNameAlreadyExists represents an account name already exists error
	AccountNameAlreadyExists
)

// Adapter represents an adapter
type Adapter interface {
	ToAdmin(ins Stack) (admin_stacks.Stack, error)
	ToStencil(ins Stack) (stencil_stacks.Stack, error)
}

// Builder represents the stack builder
type Builder interface {
	Create() Builder
	WithList(list []Frame) Builder
	Now() (Stack, error)
}

// Stack represents a stack
type Stack interface {
	Hash() hash.Hash
	Memory() Memory
	List() []Frame
	Body() []Frame // returns al the frames except the last one
	Last() Frame
	ContainsError() bool
}

// MemoryBuilder represents a memory builder
type MemoryBuilder interface {
	Create() MemoryBuilder
	WithAuthorized(authorized admin_accounts.Account) MemoryBuilder
	WithAuthenticated(authenticated identity_accounts.Account) MemoryBuilder
	Now() (Memory, error)
}

// Memory represents memory
type Memory interface {
	HasAuthorized() bool
	Authorized() admin_accounts.Account
	HasAuthenticated() bool
	Authenticated() identity_accounts.Account
}

// FrameFactory represents a frame factory
type FrameFactory interface {
	Create() Frame
}

// FrameBuilder represents a frame builder
type FrameBuilder interface {
	Create() FrameBuilder
	WihtList(list []Assignment) FrameBuilder
	Now() (Frame, error)
}

// Frame represents frame
type Frame interface {
	Hash() hash.Hash
	List() []Assignment
	Fetch(name string) (Assignable, error)
	FetchUint(name string) (*uint, error)
	ContainsError() bool
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

// AssignableBuilder represents an assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithStringList(stringList []string) AssignableBuilder
	WithError(raisedError uint) AssignableBuilder
	WithAuthorize(authorize accounts.Account) AssignableBuilder
	WithCreate(create credentials.Credentials) AssignableBuilder
	WithAdmin(admin admin_stacks.Stack) AssignableBuilder
	WithBytes(bytes []byte) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	Hash() hash.Hash
	IsStringList() bool
	StringList() []string
	IsError() bool
	Error() *uint
	IsAuthorize() bool
	Authorize() accounts.Account
	IsCreate() bool
	Create() credentials.Credentials
	IsAdmin() bool
	Admin() admin_stacks.Stack
	IsBytes() bool
	Bytes() []byte
}
