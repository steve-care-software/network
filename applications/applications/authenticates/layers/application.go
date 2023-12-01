package layers

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts/commands/layers"
)

type application struct {
	repository layers.Repository
	service    layers.Service
}

func createApplication(
	repository layers.Repository,
	service layers.Service,
) Application {
	out := application{
		repository: repository,
		service:    service,
	}

	return &out
}

// Amount returns the amount of layers
func (app *application) Amount() (uint, error) {
	return app.repository.Amount()
}

// List returns the list of layer hashes
func (app *application) List(index uint, amount uint) ([]hash.Hash, error) {
	return app.repository.List(index, amount)
}

// Exists returns true if the layer exists, false otherwise
func (app *application) Exists(hash hash.Hash) (bool, error) {
	return app.repository.Exists(hash)
}

// Retrieve retrieves a layer by hash
func (app *application) Retrieve(hash hash.Hash) (layers.Layer, error) {
	return app.repository.Retrieve(hash)
}

// Insert inserts a layer
func (app *application) Insert(layer layers.Layer) error {
	return app.service.Insert(layer)
}

// Delete deletes a layer by hash
func (app *application) Delete(hash hash.Hash) error {
	return app.service.Delete(hash)
}
