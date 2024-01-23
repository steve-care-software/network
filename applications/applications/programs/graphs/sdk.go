package graphs

import (
	"steve.care/network/applications/applications/programs/blocks"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs"
)

// NewApplication creates a new application
func NewApplication(
	blockApplication blocks.Application,
	repository programs.Repository,
	service programs.Service,
) Application {
	builder := programs.NewBuilder()
	metadataBuilder := programs.NewMetaDataBuilder()
	return createApplication(
		blockApplication,
		repository,
		service,
		builder,
		metadataBuilder,
	)
}

// Application represents the graph application
type Application interface {
	// List returns the list of all program hashes at their latest height
	List(pIsActive *bool) ([]hash.Hash, error)

	// Children returns the children's program name by its path
	Children(path []string, pIsActive *bool) ([]string, error)

	// Retrieve returns the program by hash
	Retrieve(hash hash.Hash) (programs.Program, error)

	// Revision returns the program based on its path and height
	Revision(path []string, height uint) (programs.Program, error)

	// Current returns the program based on its path at the latest height
	Current(path []string) (programs.Program, error)

	// Insert inserts a new program
	Insert(path []string, description string) error

	// Rewind rewinds the program
	Rewind(path []string) error

	// Delete deletes a program
	Delete(hash hash.Hash) error
}
