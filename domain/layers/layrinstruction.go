package layers

import (
	"steve.care/network/domain/hash"
)

type layerInstruction struct {
	hash   hash.Hash
	save   Layer
	delete hash.Hash
}

func createLayerInstructionWithSave(
	hash hash.Hash,
	save Layer,
) LayerInstruction {
	return createLayerInstructionInternally(hash, save, nil)
}

func createLayerInstructionWithDelete(
	hash hash.Hash,
	delete hash.Hash,
) LayerInstruction {
	return createLayerInstructionInternally(hash, nil, delete)
}

func createLayerInstructionInternally(
	hash hash.Hash,
	save Layer,
	delete hash.Hash,
) LayerInstruction {
	out := layerInstruction{
		hash:   hash,
		save:   save,
		delete: delete,
	}

	return &out
}

// Hash returns the hash
func (obj *layerInstruction) Hash() hash.Hash {
	return obj.hash
}

// IsSave returns true if save, false otherwise
func (obj *layerInstruction) IsSave() bool {
	return obj.save != nil
}

// Save returns save, if any
func (obj *layerInstruction) Save() Layer {
	return obj.save
}

// IsDelete returns true if delete, false otherwise
func (obj *layerInstruction) IsDelete() bool {
	return obj.delete != nil
}

// Delete returns delete, if any
func (obj *layerInstruction) Delete() hash.Hash {
	return obj.delete
}
