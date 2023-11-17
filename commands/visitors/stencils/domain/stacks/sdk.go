package stacks

import (
	"steve.care/network/commands/visitors/stencils/domain/layers"
)

// Builder represents a stack builder
type Builder interface {
	Create() Builder
	WithFrames(frames Frames) Builder
	Now() (Stack, error)
}

// Stack represents a stack
type Stack interface {
	Frames() Frames
	Body() Frames
	Last() Frame
	HasInstructions() bool
	Instructions() Instructions
}

// FramesBuilder represents the frames builder
type FramesBuilder interface {
	Create() FramesBuilder
	WithList(list []Frame) FramesBuilder
	Now() (Frames, error)
}

// Frames represents frames
type Frames interface {
	List() []Frame
}

// FrameBuilder represents the frame builder
type FrameBuilder interface {
	Create() FrameBuilder
	WithInstructions(instructions Instructions) FrameBuilder
	WithAssignments(assignments Assignments) FrameBuilder
	Now() (Frame, error)
}

// Frame represents a frame
type Frame interface {
	HasInstructions() bool
	Instructions() Instructions
	HasInput() bool
	Input() []byte
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// Instruction represents an instruction
type Instruction interface {
	IsSave() bool
	Save() Resource
}

// Resource represents a resource
type Resource interface {
	Path() []string
	Layer() layers.Layer
}

// Assignments represents assignments
type Assignments interface {
	List() []Assignment
}

// Assignment represents an assignment
type Assignment interface {
	Name() string
	Assignable() Assignable
}

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithBytes(bytes []byte) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsBytes() bool
	Bytes() []byte
}
