package programs

import "steve.care/network/domain/programs"

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
	IsVisitor() bool
	Visitor() programs.Program
}
