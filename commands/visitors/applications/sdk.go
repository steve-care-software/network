package applications

import (
	"steve.care/network/commands/visitors/domain/programs"
	"steve.care/network/libraries/stacks"
)

// Application represents the application
type Application interface {
	Execute(program programs.Program, stack stacks.Stack) (stacks.Stack, error)
}
