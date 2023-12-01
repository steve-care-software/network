package sqllites

import (
	"steve.care/network/applications/applications/encryptors"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/credentials"
)

type accountService struct {
	encryptorApp encryptors.Application
	adapter      accounts.Adapter
}

func createAccountService(
	encryptorApp encryptors.Application,
	adapter accounts.Adapter,
) accounts.Service {
	out := accountService{
		encryptorApp: encryptorApp,
		adapter:      adapter,
	}

	return &out
}

// Insert inserts an account
func (app *accountService) Insert(account accounts.Account, password []byte) error {
	return nil
}

// Update updates an account
func (app *accountService) Update(credentials credentials.Credentials, criteria accounts.UpdateCriteria) error {
	return nil
}

// Delete deletes an account
func (app *accountService) Delete(credentials credentials.Credentials) error {
	return nil
}
