package programs

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/logics"
)

type program struct {
	hash        hash.Hash
	name        string
	description string
	head        blocks.Block
	logic       logics.Logic
	parent      Program
}

func createProgram(
	hash hash.Hash,
	name string,
	description string,
) Program {
	return createProgramInternally(hash, name, description, nil, nil, nil)
}

func createProgramWithHead(
	hash hash.Hash,
	name string,
	description string,
	head blocks.Block,
) Program {
	return createProgramInternally(hash, name, description, head, nil, nil)
}

func createProgramWithLogic(
	hash hash.Hash,
	name string,
	description string,
	logic logics.Logic,
) Program {
	return createProgramInternally(hash, name, description, nil, logic, nil)
}

func createProgramWithParent(
	hash hash.Hash,
	name string,
	description string,
	parent Program,
) Program {
	return createProgramInternally(hash, name, description, nil, nil, parent)
}

func createProgramWithHeadAndLogic(
	hash hash.Hash,
	name string,
	description string,
	head blocks.Block,
	logic logics.Logic,
) Program {
	return createProgramInternally(hash, name, description, head, logic, nil)
}

func createProgramWithHeadAndParent(
	hash hash.Hash,
	name string,
	description string,
	head blocks.Block,
	parent Program,
) Program {
	return createProgramInternally(hash, name, description, head, nil, parent)
}

func createProgramWithLogicAndParent(
	hash hash.Hash,
	name string,
	description string,
	logic logics.Logic,
	parent Program,
) Program {
	return createProgramInternally(hash, name, description, nil, logic, nil)
}

func createProgramWithHeadAndLogicAndParent(
	hash hash.Hash,
	name string,
	description string,
	head blocks.Block,
	logic logics.Logic,
	parent Program,
) Program {
	return createProgramInternally(hash, name, description, head, logic, parent)
}

func createProgramInternally(
	hash hash.Hash,
	name string,
	description string,
	head blocks.Block,
	logic logics.Logic,
	parent Program,
) Program {
	out := program{
		hash:        hash,
		name:        name,
		description: description,
		head:        head,
		logic:       logic,
		parent:      parent,
	}

	return &out
}

// Hash returns the hash
func (obj *program) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *program) Name() string {
	return obj.name
}

// Description returns the description
func (obj *program) Description() string {
	return obj.description
}

// HasHead returns true if there is a head, false otherwise
func (obj *program) HasHead() bool {
	return obj.head != nil
}

// Head returns the head, if any
func (obj *program) Head() blocks.Block {
	return obj.head
}

// HasLogic returns true if there is a logic, false otherwise
func (obj *program) HasLogic() bool {
	return obj.logic != nil
}

// Logic returns the logic, if any
func (obj *program) Logic() logics.Logic {
	return obj.logic
}

// HasParent returns true if there is a parent, false otherwise
func (obj *program) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *program) Parent() Program {
	return obj.parent
}
