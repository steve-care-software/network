package resources

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/executions/actions/resources"
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

// AmountByQuery returns the amount of resources by criteria
func (app *application) AmountByQuery(criteria hash.Hash) (uint, error) {
	return app.repository.AmountByQuery(criteria)
}

// ListByQuery lists resource hashes by criteria
func (app *application) ListByQuery(criteria hash.Hash) ([]hash.Hash, error) {
	return app.repository.ListByQuery(criteria)
}

// RetrieveByQuery retrieves a resource by criteria
func (app *application) RetrieveByQuery(criteria hash.Hash) (resources.Resource, error) {
	return app.repository.RetrieveByQuery(criteria)
}

// RetrieveByHash retrieves a resource by hash
func (app *application) RetrieveByHash(hash hash.Hash) (resources.Resource, error) {
	return app.repository.RetrieveByHash(hash)
}

// Insert inserts a resource
func (app *application) Insert(ins resources.Resource) error {
	return app.service.Insert(ins)
}

// Delete deletes a resource
func (app *application) Delete(hash hash.Hash) error {
	return app.service.Delete(hash)
}
