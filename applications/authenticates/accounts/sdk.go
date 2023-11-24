package accounts

import "steve.care/network/domain/accounts"

// Builder represents the account application
type Builder interface {
	Create() Builder
	WithAccount(account accounts.Account) Builder
	Now() (Application, error)
}

// Application represents the authenticated account application
type Application interface {
	Retrieve() (accounts.Account, error)
	Delete(password []byte) error
	Update(currentPassword []byte, newPassword []byte) error
}
