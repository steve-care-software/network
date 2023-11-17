package programs

import "steve.care/network/commands/visitors/admins/identities/dashboards/stencils/layers/domain/layers"

// Adapter represents the program adapter
type Adapter interface {
	ToBytes(ins Program) ([]byte, error)
	ToInstance(bytes []byte) (Program, error)
}

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
	IsAssignment() bool
	Assignment() Assignment
	IsSave() bool
	Save() Save
	IsMove() bool
	Move() Move
	IsDelete() bool
	Delete() []string
}

// Move represents a move layer
type Move interface {
	From() []string
	To() []string
}

// Save represents a save layer
type Save interface {
	Path() []string
	Layer() layers.Layer
}

// Assignment represents an assignment
type Assignment interface {
	Name() string
	Assignable() Assignable
}

// Assignable represents an assignable
type Assignable interface {
	IsExists() bool
	Exists() []string
	IsList() bool
	List() []string
	IsDir() bool
	Dir() []string
	IsRetrieve() bool
	Retrieve() []string
}
