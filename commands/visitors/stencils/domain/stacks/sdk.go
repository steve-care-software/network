package stacks

import (
	admin_accounts "steve.care/network/commands/visitors/admins/domain/accounts"
	identity_accounts "steve.care/network/commands/visitors/admins/identities/domain/accounts"
)

// Stack represents a stack
type Stack interface {
	Memory() Memory
}

// Memory represents the memory
type Memory interface {
	HasAuthorized() bool
	Authorized() admin_accounts.Account
	HasAuthenticated() bool
	Authenticated() identity_accounts.Account
}
