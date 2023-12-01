package sqllites

import (
	"steve.care/network/applications/applications/encryptors"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/credentials"
)

type accountRepository struct {
	encryptorApp encryptors.Application
	adapter      accounts.Adapter
}

func createAccountRepository(
	encryptorApp encryptors.Application,
	adapter accounts.Adapter,
) accounts.Repository {
	out := accountRepository{
		encryptorApp: encryptorApp,
		adapter:      adapter,
	}

	return &out
}

// List returns the list of usernames
func (app *accountRepository) List() ([]string, error) {
	return nil, nil
}

// Exists returns true if the username exists, false otherwise
func (app *accountRepository) Exists(username string) (bool, error) {
	return false, nil
}

// Retrieve retrieves the accont from credentials
func (app *accountRepository) Retrieve(credentials credentials.Credentials) (accounts.Account, error) {
	return nil, nil
}
