package programs

// Program represents a visitor's program
type Program interface {
	Instructions() Instructions
}

// Instructions represents an instructions
type Instructions interface {
	List() []Instruction
}

// Instruction represents an instruction
type Instruction interface {
	IsAssignment() bool
	Assignment() Assignment
}

// Assignment represents an assignment
type Assignment interface {
	Name() string
	Assignable() Assignable
}

// Assignable represents an assignable
type Assignable interface {
	IsAuthenticate() bool
	Authenticate() Credentials
	IsCreateAdministrator() bool
	CreateAdministrator() Credentials
}

// Credentials represents credentials
type Credentials interface {
	Username() string
	Password() []byte
}
