package accounts

import (
	"steve.care/network/domain/accounts"
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
)

// NewApplication creates a new application
func NewApplication(
	repository accounts.Repository,
	service accounts.Service,
	bitrate int,
) Application {
	accountBuilder := accounts.NewBuilder()
	signerFactory := signers.NewFactory()
	encryptorBuilder := account_encryptors.NewBuilder()

	return createApplication(
		accountBuilder,
		signerFactory,
		encryptorBuilder,
		repository,
		service,
		bitrate,
	)
}

// Application represents the authenticated account application
type Application interface {
	List() ([]string, error)
	Exists(username string) (bool, error)
	Insert(credentials credentials.Credentials) error
	Retrieve(credentials credentials.Credentials) (accounts.Account, error)
	Update(credentials credentials.Credentials, criteria accounts.UpdateCriteria) error
	Delete(credentials credentials.Credentials) error
}
