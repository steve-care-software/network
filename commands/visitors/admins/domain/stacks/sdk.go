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
	WithList(list []Frame) Builder
	Now() (Stack, error)
}

// Stack represents a stack
type Stack interface {
	Hash() hash.Hash
	List() []Frame
	Body() []Frame
	Last() Frame
	Authorized() accounts.Account
	ContainsError() bool
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

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithBool(boolValue bool) AssignableBuilder
	WithAuthorized(authorized accounts.Account) AssignableBuilder
	WithIdentities(identities accounts.Identities) AssignableBuilder
	WithCreateAccount(createAccount CreateAccount) AssignableBuilder
	WithError(errorCode uint) AssignableBuilder
	Now() (Assignable, error)
}

type Assignable interface {
	Hash() hash.Hash
	IsBool() bool
	Bool() *bool
	IsAuthorized() bool
	Authorized() accounts.Account
	IsIdentities() bool
	Identities() accounts.Identities
	IsCreateAccount() bool
	CreateAccount() CreateAccount
	IsError() bool
	Error() uint
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
