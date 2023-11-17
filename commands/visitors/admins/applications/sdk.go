package applications

import (
	"steve.care/network/commands/visitors/admins/domain/programs"
	"steve.care/network/commands/visitors/admins/domain/stacks"
)

// Application represents the admin application
type Application interface {
	ExecuteBytes(bytes []byte, stack stacks.Stack) (stacks.Stack, error)
	Execute(program programs.Program, stack stacks.Stack) (stacks.Stack, error)
}
