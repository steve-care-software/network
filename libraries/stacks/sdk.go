package stacks

import "steve.care/network/libraries/blockchains/blocks/queues"

// Builder represents the stack builder
type Builder interface {
	Create() Builder
	WithList(list []Frame) Builder
	Now() (Stack, error)
}

// Stack represents a stack
type Stack interface {
	List() []Frame
	Last() Frame
}

// FrameFactory represents a frame factory
type FrameFactory interface {
	Create() Frame
}

// FrameBuilder represents a frame builder
type FrameBuilder interface {
	Create() FrameBuilder
	WihtList(list []Assignable) FrameBuilder
	Now() (Frame, error)
}

// Frame represents frame
type Frame interface {
	List() []Assignable
	Fetch(name string) (Assignable, error)
	FetchUint(name string) (*uint, error)
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
	WithUint(uintValue uint) AssignableBuilder
	WithQueue(queue queues.Queue) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsUint() bool
	Uint() *uint
	IsQueue() bool
	Queue() queues.Queue
}
