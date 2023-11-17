package programs

// Program represents a program
type Program interface {
	Instructions() Instructions
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// Instruction represents an instruction
type Instruction interface {
	IsDeleteAuthorized() bool
	IsCreateadmin() bool
	CreateAdmin() Credentials
}

// Credentials represents credentials
type Credentials interface {
	Username() string
	Password() []byte
}

// Assignment represents an assignment
type Assignment interface {
	Name() string
	Assignable() Assignable
}

// Assignable represents an assignable
type Assignable interface {
	IsListIdentities() bool
}
