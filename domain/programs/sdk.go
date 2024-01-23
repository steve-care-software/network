package programs

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks"
	"steve.care/network/domain/programs/logics"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewMetaDataBuilder creates a new metadata builder
func NewMetaDataBuilder() MetaDataBuilder {
	hashAdapter := hash.NewAdapter()
	return createMetaDataBuilder(
		hashAdapter,
	)
}

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithDescription(description string) Builder
	WithLogic(logic logics.Logic) Builder
	WithHead(head blocks.Block) Builder
	WithMetaData(metadata MetaData) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Description() string
	HasHead() bool
	Head() blocks.Block
	HasLogic() bool
	Logic() logics.Logic
	HasMetaData() bool
	MetaData() MetaData
}

// MetaDataBuilder represents metadata builder
type MetaDataBuilder interface {
	Create() MetaDataBuilder
	WithName(name string) MetaDataBuilder
	WithParent(parent Program) MetaDataBuilder
	Now() (MetaData, error)
}

// MetaData represents metadata
type MetaData interface {
	Hash() hash.Hash
	Name() string
	Parent() Program
}

// Repository represents the program repository
type Repository interface {
	// List returns the list of all program hashes at their latest height
	List(isActive *bool) ([]hash.Hash, error)

	// Children returns the children's program paths by its path
	Children(path []string, pIsActive *bool) ([]string, error)

	// Retrieve returns the program by hash
	Retrieve(hash hash.Hash) (Program, error)

	// Revision returns the program based on its path and height
	Revision(path []string, height uint) (Program, error)

	// Current returns the program based on its path at the latest height
	Current(path []string) (Program, error)
}

// Service represents the program service
type Service interface {
	// Insert inserts a program
	Insert(program Program) error

	// Delete deletes a program
	Delete(program Program) error
}
