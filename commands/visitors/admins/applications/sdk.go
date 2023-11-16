package applications

import (
	"steve.care/network/commands/visitors/admins/domain/instructions"
	"steve.care/network/commands/visitors/admins/domain/stacks"
)

// Application represents the admin application
type Application interface {
	Execute(instruction instructions.Instruction, stack stacks.Stack) (stacks.Stack, error)
}
