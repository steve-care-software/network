package applications

import (
	"steve.care/network/commands/visitors/domain/programs"
	"steve.care/network/commands/visitors/domain/stacks"
)

// Application represents the application
type Application interface {
	ExecuteBytes(bytes []byte, stack stacks.Stack) (stacks.Stack, error)
	Execute(program programs.Program, stack stacks.Stack) (stacks.Stack, error)
}
