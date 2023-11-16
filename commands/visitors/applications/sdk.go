package applications

import (
	"steve.care/8web/commands/visitors/domain/programs"
	"steve.care/8web/commands/visitors/domain/stacks"
)

// Application represents the application
type Application interface {
	Execute(program programs.Program, stack stacks.Stack) (stacks.Stack, error)
}
