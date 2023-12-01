package receipts

import (
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/receipts"
)

type application struct {
	repository receipts.Repository
	service    receipts.Service
}

func createApplication(
	repository receipts.Repository,
	service receipts.Service,
) Application {
	out := application{
		repository: repository,
		service:    service,
	}

	return &out
}

// Amount returns the amount of receipts
func (app *application) Amount() (uint, error) {
	return app.repository.Amount()
}

// List returns the list of receipts
func (app *application) List(index uint, amount uint) ([]hash.Hash, error) {
	return app.repository.List(index, amount)
}

// ListBySigner returns the list of receipts by signer
func (app *application) ListBySigner(pubKey signers.PublicKey, index uint, amount uint) ([]hash.Hash, error) {
	return app.repository.ListBySigner(pubKey, index, amount)
}

// Retrieve retrieves a receipt
func (app *application) Retrieve(hash hash.Hash) (receipts.Receipt, error) {
	return app.repository.Retrieve(hash)
}

// Insert inserts a receipt
func (app *application) Insert(receipt receipts.Receipt) error {
	return app.service.Insert(receipt)
}

// Delete deletes a receipt by hash
func (app *application) Delete(hash hash.Hash) error {
	return app.service.Delete(hash)
}
