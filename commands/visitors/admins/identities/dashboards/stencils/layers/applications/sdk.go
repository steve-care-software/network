package applications

import (
	"steve.care/network/commands/visitors/admins/identities/dashboards/stencils/layers/domain/programs"
	"steve.care/network/commands/visitors/admins/identities/dashboards/stencils/layers/domain/stacks"
)

// Application represents a layer application
type Application interface {
	ExecuteBytes(bytes []byte, stack stacks.Stack) (stacks.Stack, error)
	Execute(program programs.Program, stack stacks.Stack) (stacks.Stack, error)
	Process(context uint, stack stacks.Stack) error
}
