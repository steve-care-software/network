package layers

import "steve.care/network/domain/hash"

type instructionResource struct {
	hash hash.Hash
	save string
	del  string
}

func createInstructionResourceWithSave(
	hash hash.Hash,
	save string,
) InstructionResource {
	return createInstructionResourceInternally(hash, save, "")
}

func createInstructionResourceWithDelete(
	hash hash.Hash,
	del string,
) InstructionResource {
	return createInstructionResourceInternally(hash, "", del)
}

func createInstructionResourceInternally(
	hash hash.Hash,
	save string,
	del string,
) InstructionResource {
	out := instructionResource{
		hash: hash,
		save: save,
		del:  del,
	}

	return &out
}

// Hash returns the hash
func (obj *instructionResource) Hash() hash.Hash {
	return obj.hash
}

// IsSave returns true if save, false otherwise
func (obj *instructionResource) IsSave() bool {
	return obj.save != ""
}

// Save returns the save, if any
func (obj *instructionResource) Save() string {
	return obj.save
}

// IsDelete returns true if delete, false otherwise
func (obj *instructionResource) IsDelete() bool {
	return obj.del != ""
}

// Delete returns the delete, if any
func (obj *instructionResource) Delete() string {
	return obj.del
}
