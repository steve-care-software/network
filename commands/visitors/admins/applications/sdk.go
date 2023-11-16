package applications

import (
	"steve.care/network/commands/visitors/admins/domain/instructions"
	"steve.care/network/commands/visitors/admins/domain/stacks"
	"steve.care/network/libraries/credentials"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the admin application
type Application interface {
	Execute(instruction instructions.Instruction, stack stacks.Stack) (stacks.Stack, error)
}
