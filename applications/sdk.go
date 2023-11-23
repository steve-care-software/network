package applications

import (
	identity_accounts "steve.care/network/domain/accounts"
	"steve.care/network/domain/results"
	"steve.care/network/domain/stacks"
)

// Application represents a stencil application
type Application interface {
	Execute(authenticated identity_accounts.Account, stack stacks.Stack) (results.Result, error)
}
