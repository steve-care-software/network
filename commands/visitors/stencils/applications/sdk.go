package applications

import (
	admin_accounts "steve.care/network/commands/visitors/admins/domain/accounts"
	identity_accounts "steve.care/network/commands/visitors/admins/identities/domain/accounts"
	"steve.care/network/commands/visitors/stencils/domain/stacks"
)

// Application represents a stencil application
type Application interface {
	Execute(authorized admin_accounts.Account, authenticated identity_accounts.Account, stack stacks.Stack) ([]byte, error)
}
