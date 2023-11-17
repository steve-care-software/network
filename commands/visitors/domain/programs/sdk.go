package programs

import (
	admin_programs "steve.care/network/commands/visitors/admins/domain/programs"
	stencils_program "steve.care/network/commands/visitors/stencils/domain/programs"
)

// Adapter represents a program adapter
type Adapter interface {
	ToBytes(ins Program) ([]byte, error)
	ToInstance(bytes []byte) (Program, error)
}

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithInstructions(instructions Instructions) Builder
	Now() (Program, error)
}

// Program represents a visitor's program
type Program interface {
	Instructions() Instructions
}

// InstructionsBuilder represents an instructions builder
type InstructionsBuilder interface {
	Create() InstructionsBuilder
	WithList(list []Instruction) InstructionsBuilder
	Now() (Instructions, error)
}

// Instructions represents an instructions
type Instructions interface {
	List() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithAssignment(assignment Assignment) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsAssignment() bool
	Assignment() Assignment
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithName(name string) AssignmentBuilder
	WithAssignable(assignable Assignable) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Name() string
	Assignable() Assignable
}

// AssignableBuilder represents an assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithAuthorize(authorize Credentials) AssignableBuilder
	WithCreate(create Credentials) AssignableBuilder
	WithStencil(stencil stencils_program.Program) AssignableBuilder
	WithAdmin(admin admin_programs.Program) AssignableBuilder
	IsListNames() AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsListNames() bool
	IsAuthorize() bool
	Authorize() Credentials
	IsCreate() bool
	Create() Credentials
	IsStencil() bool
	Stencil() stencils_program.Program
	IsAdmin() bool
	Admin() admin_programs.Program
}

// CredentialsBuilder represents a credentials builer
type CredentialsBuilder interface {
	Create() CredentialsBuilder
	WithUsername(username string) CredentialsBuilder
	WithPassword(password []byte) CredentialsBuilder
	Now() (Credentials, error)
}

// Credentials represents credentials
type Credentials interface {
	Username() string
	Password() []byte
}
