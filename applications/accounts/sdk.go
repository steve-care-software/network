package accounts

import (
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/credentials"
)

// Application represents the authenticated account application
type Application interface {
	List() ([]string, error)
	Exists(username string) (bool, error)
	Insert(credentials credentials.Credentials) error
	Retrieve(credentials credentials.Credentials) (accounts.Account, error)
	Update(credentials credentials.Credentials, criteria accounts.UpdateCriteria) error
	Delete(credentials credentials.Credentials) error
}
