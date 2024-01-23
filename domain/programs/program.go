package programs

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/logics"
)

type program struct {
	hash        hash.Hash
	description string
	head        blocks.Block
	logic       logics.Logic
	metaData    MetaData
}

func createProgram(
	hash hash.Hash,
	description string,
) Program {
	return createProgramInternally(hash, description, nil, nil, nil)
}

func createProgramWithHead(
	hash hash.Hash,
	description string,
	head blocks.Block,
) Program {
	return createProgramInternally(hash, description, head, nil, nil)
}

func createProgramWithLogic(
	hash hash.Hash,
	description string,
	logic logics.Logic,
) Program {
	return createProgramInternally(hash, description, nil, logic, nil)
}

func createProgramWithMetaData(
	hash hash.Hash,
	description string,
	metaData MetaData,
) Program {
	return createProgramInternally(hash, description, nil, nil, metaData)
}

func createProgramWithHeadAndLogic(
	hash hash.Hash,
	description string,
	head blocks.Block,
	logic logics.Logic,
) Program {
	return createProgramInternally(hash, description, head, logic, nil)
}

func createProgramWithHeadAndMetaData(
	hash hash.Hash,
	description string,
	head blocks.Block,
	metaData MetaData,
) Program {
	return createProgramInternally(hash, description, head, nil, metaData)
}

func createProgramWithLogicAndMetaData(
	hash hash.Hash,
	description string,
	logic logics.Logic,
	metaData MetaData,
) Program {
	return createProgramInternally(hash, description, nil, logic, metaData)
}

func createProgramWithHeadAndLogicAndMetaData(
	hash hash.Hash,
	description string,
	head blocks.Block,
	logic logics.Logic,
	metaData MetaData,
) Program {
	return createProgramInternally(hash, description, head, logic, metaData)
}

func createProgramInternally(
	hash hash.Hash,
	description string,
	head blocks.Block,
	logic logics.Logic,
	metaData MetaData,
) Program {
	out := program{
		hash:        hash,
		description: description,
		head:        head,
		logic:       logic,
		metaData:    metaData,
	}

	return &out
}

// Hash returns the hash
func (obj *program) Hash() hash.Hash {
	return obj.hash
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

// HasMetaData returns true if there is metadata, false otherwise
func (obj *program) HasMetaData() bool {
	return obj.metaData != nil
}

// MetaData returns the metadata, if any
func (obj *program) MetaData() MetaData {
	return obj.metaData
}
