package applications

import (
	"steve.care/network/domain/programs"
	"steve.care/network/domain/stacks"
)

// Application represents an application
type Application interface {
	Execute(programm programs.Program, stack stacks.Stack) (stacks.Stack, error)
}
