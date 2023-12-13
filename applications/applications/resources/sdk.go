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

	// AmountByCriteria returns the amount of resources by criteria
	AmountByCriteria(criteria hash.Hash) (uint, error)

	// ListByCriteria lists resource hashes by criteria
	ListByCriteria(criteria hash.Hash) ([]hash.Hash, error)

	// RetrieveByCriteria retrieves a resource by criteria
	RetrieveByCriteria(criteria hash.Hash) (resources.Resource, error)

	// Insert inserts a resource
	Insert(ins resources.Resource) error

	// Delete deletes a resource
	Delete(hash hash.Hash) error
}
