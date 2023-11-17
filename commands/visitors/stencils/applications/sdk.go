package applications

import (
	"steve.care/network/commands/visitors/stencils/domain/programs"
	"steve.care/network/commands/visitors/stencils/domain/results"
	"steve.care/network/commands/visitors/stencils/domain/stacks"
)

// Application represents a stencil application
type Application interface {
	Execute(program programs.Program, stack stacks.Stack) (results.Result, error)
}
