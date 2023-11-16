package applications

import (
	"steve.care/network/commands/domain/commands"
	"steve.care/network/commands/visitors/domain/programs"
	"steve.care/network/commands/visitors/domain/stacks"
)

// Application represents the commands application
type Application interface {
	ExecuteBytes(bytes []byte, stack stacks.Stack) (commands.Command, error)
	Execute(program programs.Program, stack stacks.Stack) (commands.Command, error)
}
