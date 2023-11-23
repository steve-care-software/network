package authenticates

import (
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/commands/links"
	"steve.care/network/domain/commands/results"
	"steve.care/network/domain/hash"
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
	Execute(hash hash.Hash, input []byte) (results.Result, error)
	Links(executed []hash.Hash) (links.Link, error)
	Clear() error
}
