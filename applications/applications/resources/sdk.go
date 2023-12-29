package resources

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/resources"
)

// NewApplication creates a new application
func NewApplication(
	repository resources.Repository,
	service resources.Service,
) Application {
	return createApplication(
		repository,
		service,
	)
}

// Application represents the resources application
type Application interface {
	// Amount returns the amount of resources
	Amount() (uint, error)

	// AmountByQuery returns the amount of resources by criteria
	AmountByQuery(criteria hash.Hash) (uint, error)

	// ListByQuery lists resource hashes by criteria
	ListByQuery(criteria hash.Hash) ([]hash.Hash, error)

	// RetrieveByQuery retrieves a resource by criteria
	RetrieveByQuery(criteria hash.Hash) (resources.Resource, error)

	// RetrieveByHash retrieves a resource by hash
	RetrieveByHash(hash hash.Hash) (resources.Resource, error)

	// Insert inserts a resource
	Insert(ins resources.Resource) error

	// Delete deletes a resource
	Delete(hash hash.Hash) error
}
