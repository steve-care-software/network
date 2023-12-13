package resources

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/resources"
)

type application struct {
	repository resources.Repository
	service    resources.Service
}

func createApplication(
	repository resources.Repository,
	service resources.Service,
) Application {
	out := application{
		repository: repository,
		service:    service,
	}

	return &out
}

// Amount returns the amount of resources
func (app *application) Amount() (uint, error) {
	return app.repository.Amount()
}

// AmountByCriteria returns the amount of resources by criteria
func (app *application) AmountByCriteria(criteria hash.Hash) (uint, error) {
	return app.repository.AmountByCriteria(criteria)
}

// ListByCriteria lists resource hashes by criteria
func (app *application) ListByCriteria(criteria hash.Hash) ([]hash.Hash, error) {
	return app.repository.ListByCriteria(criteria)
}

// RetrieveByCriteria retrieves a resource by criteria
func (app *application) RetrieveByCriteria(criteria hash.Hash) (resources.Resource, error) {
	return app.repository.RetrieveByCriteria(criteria)
}

// Insert inserts a resource
func (app *application) Insert(ins resources.Resource) error {
	return app.service.Insert(ins)
}

// Delete deletes a resource
func (app *application) Delete(hash hash.Hash) error {
	return app.service.Delete(hash)
}
