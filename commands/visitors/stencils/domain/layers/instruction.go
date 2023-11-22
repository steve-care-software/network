package layers

import "steve.care/network/libraries/hash"

type instruction struct {
	hash       hash.Hash
	isStop     bool
	raiseError uint
	condition  Condition
	save       Layer
	assignment Assignment
}

func createInstructionWithIsStop(
	hash hash.Hash,
) Instruction {
	return createInstructionInternally(
		hash,
		true,
		0,
		nil,
		nil,
		nil,
	)
}

func createInstructionWithRaiseError(
	hash hash.Hash,
	raiseError uint,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		raiseError,
		nil,
		nil,
		nil,
	)
}

func createInstructionWithCondition(
	hash hash.Hash,
	condition Condition,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		0,
		condition,
		nil,
		nil,
	)
}

func createInstructionWithSave(
	hash hash.Hash,
	save Layer,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		0,
		nil,
		save,
		nil,
	)
}

func createInstructionWithAssignment(
	hash hash.Hash,
	assignment Assignment,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		0,
		nil,
		nil,
		assignment,
	)
}

func createInstructionInternally(
	hash hash.Hash,
	isStop bool,
	raiseError uint,
	condition Condition,
	save Layer,
	assignment Assignment,
) Instruction {
	out := instruction{
		hash:       hash,
		isStop:     isStop,
		raiseError: raiseError,
		condition:  condition,
		save:       save,
		assignment: assignment,
	}

	return &out
}

// Hash returns the hash
func (obj *instruction) Hash() hash.Hash {
	return obj.hash
}

// IsStop returns true if stop, false otherwise
func (obj *instruction) IsStop() bool {
	return obj.isStop
}

// IsRaiseError returns true if raiseError, false otherwise
func (obj *instruction) IsRaiseError() bool {
	return obj.raiseError > 0
}

// RaiseError returns the raiseError, if any
func (obj *instruction) RaiseError() uint {
	return obj.raiseError
}

// IsCondition returns true if condition, false otherwise
func (obj *instruction) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *instruction) Condition() Condition {
	return obj.condition
}

// IsSave returns true if save, false otherwise
func (obj *instruction) IsSave() bool {
	return obj.save != nil
}

// Save returns the save, if any
func (obj *instruction) Save() Layer {
	return obj.save
}

// IsAssignment returns true if assignment, false otherwise
func (obj *instruction) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *instruction) Assignment() Assignment {
	return obj.assignment
}
