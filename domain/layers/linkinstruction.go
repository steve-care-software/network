package layers

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/links"
)

type linkInstruction struct {
	hash   hash.Hash
	save   links.Link
	delete hash.Hash
}

func createLinkInstructionWithSave(
	hash hash.Hash,
	save links.Link,
) LinkInstruction {
	return createLinkInstructionInternally(hash, save, nil)
}

func createLinkInstructionWithDelete(
	hash hash.Hash,
	delete hash.Hash,
) LinkInstruction {
	return createLinkInstructionInternally(hash, nil, delete)
}

func createLinkInstructionInternally(
	hash hash.Hash,
	save links.Link,
	delete hash.Hash,
) LinkInstruction {
	out := linkInstruction{
		hash:   hash,
		save:   save,
		delete: delete,
	}

	return &out
}

// Hash returns the hash
func (obj *linkInstruction) Hash() hash.Hash {
	return obj.hash
}

// IsSave returns true if save, false otherwise
func (obj *linkInstruction) IsSave() bool {
	return obj.save != nil
}

// Save returns save, if any
func (obj *linkInstruction) Save() links.Link {
	return obj.save
}

// IsDelete returns true if delete, false otherwise
func (obj *linkInstruction) IsDelete() bool {
	return obj.delete != nil
}

// Delete returns delete, if any
func (obj *linkInstruction) Delete() hash.Hash {
	return obj.delete
}
