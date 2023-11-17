package stacks

import "steve.care/network/commands/visitors/admins/domain/accounts"

// Stack represents the stack
type Stack interface {
	List() []Frame
	Body() []Frame
	Last() []Frame
	Authorized() (accounts.Account, error)
}

// Frame represents the frame
type Frame interface {
}

type Assignable interface {
	IsAuthorized() bool
	Authorized() accounts.Account
	IsIdentities() bool
	Identities() accounts.Identities
	IsError() bool
	Error() uint
}
