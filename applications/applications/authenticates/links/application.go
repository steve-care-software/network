package links

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/links"
)

type application struct {
	repository links.Repository
	service    links.Service
}

func createApplication(
	repository links.Repository,
	service links.Service,
) Application {
	out := application{
		repository: repository,
		service:    service,
	}

	return &out
}

// Amount returns the amount of links
func (app *application) Amount() (uint, error) {
	return app.repository.Amount()
}

// List returns the list of link hashes
func (app *application) List(index uint, amount uint) ([]hash.Hash, error) {
	return app.repository.List(index, amount)
}

// Exists returns true if the link exists, false otherwise
func (app *application) Exists(hash hash.Hash) (bool, error) {
	return app.repository.Exists(hash)
}

// RetrieveByHash retrieves a link by hash
func (app *application) RetrieveByHash(hash hash.Hash) (links.Link, error) {
	return app.repository.RetrieveByHash(hash)
}

// RetrieveByOrigin retrieves a link by origin
func (app *application) RetrieveByOrigin(executedLayers []hash.Hash) (links.Link, error) {
	return app.repository.RetrieveByOrigin(executedLayers)
}

// Insert inserts a link
func (app *application) Insert(link links.Link) error {
	return app.service.Insert(link)
}

// Delete deletes a link by hash
func (app *application) Delete(hash hash.Hash) error {
	return app.service.Delete(hash)
}
