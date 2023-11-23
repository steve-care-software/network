package authenticates

import (
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/links"
	"steve.care/network/domain/results"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithAccount(account accounts.Account) Builder
	NOw() (Application, error)
}

// Application represents an authenticated application
type Application interface {
	Delete(password []byte) error
	Update(currentPassword []byte, newPassword []byte) error
	Exists(hash hash.Hash) (bool, error)
	Execute(hash hash.Hash) (results.Result, error)
	Links(executed []hash.Hash) (links.Link, error)
	Clear() error
}
