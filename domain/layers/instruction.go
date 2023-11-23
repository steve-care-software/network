package layers

import "steve.care/network/domain/hash"

type instruction struct {
	hash       hash.Hash
	isStop     bool
	raiseError uint
	condition  Condition
	assignment Assignment
	link       LinkInstruction
	layer      LayerInstruction
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
		assignment,
		nil,
		nil,
	)
}

func createInstructionWithLink(
	hash hash.Hash,
	link LinkInstruction,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		0,
		nil,
		nil,
		link,
		nil,
	)
}

func createInstructionWithLayer(
	hash hash.Hash,
	layer LayerInstruction,
) Instruction {
	return createInstructionInternally(
		hash,
		false,
		0,
		nil,
		nil,
		nil,
		layer,
	)
}

func createInstructionInternally(
	hash hash.Hash,
	isStop bool,
	raiseError uint,
	condition Condition,
	assignment Assignment,
	link LinkInstruction,
	layer LayerInstruction,
) Instruction {
	out := instruction{
		hash:       hash,
		isStop:     isStop,
		raiseError: raiseError,
		condition:  condition,
		assignment: assignment,
		link:       link,
		layer:      layer,
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

// IsAssignment returns true if assignment, false otherwise
func (obj *instruction) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *instruction) Assignment() Assignment {
	return obj.assignment
}

// IsLink returns true if link, false otherwise
func (obj *instruction) IsLink() bool {
	return obj.link != nil
}

// Link returns the link, if any
func (obj *instruction) Link() LinkInstruction {
	return obj.link
}

// IsLayer returns true if layer, false otherwise
func (obj *instruction) IsLayer() bool {
	return obj.layer != nil
}

// Layer returns the layer, if any
func (obj *instruction) Layer() LayerInstruction {
	return obj.layer
}
