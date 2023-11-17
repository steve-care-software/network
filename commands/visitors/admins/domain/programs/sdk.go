package programs

// Adapter represents the program adapter
type Adapter interface {
	ToBytes(ins Program) ([]byte, error)
	ToInstance(bytes []byte) (Program, error)
}

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithInstructions(instructions Instructions) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Instructions() Instructions
}

// InstructionsBuilder represents the instructions builder
type InstructionsBuilder interface {
	Create() InstructionsBuilder
	WithList(list []Instruction) InstructionsBuilder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// InstructionBuilder represents the instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithAssignment(assignment Assignment) InstructionBuilder
	WithCreateAdmin(createAdmin Credentials) InstructionBuilder
	IsDeleteAuthorized() InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsAssignment() bool
	Assignment() Assignment
	IsCreateAdmin() bool
	CreateAdmin() Credentials
	IsDeleteAuthorized() bool
}

// AssignmentBuilder represents the assignment builder
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

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	HasIdentities() AssignableBuilder
	ListIdentities() AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsHasIdentities() bool
	IsListIdentities() bool
}

// CredentialsBuilder represents the credentials builder
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
