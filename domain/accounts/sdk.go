package accounts

import (
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases"
	"steve.care/network/domain/encryptors"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewRepositoryBuilder creates a new reposiotry builder
func NewRepositoryBuilder(
	encryptor encryptors.Encryptor,
	adapter Adapter,
) RepositoryBuilder {
	return createRepositoryBuilder(
		encryptor,
		adapter,
	)
}

// NewServiceBuilder creates a new service builder
func NewServiceBuilder(
	encryptor encryptors.Encryptor,
	repository Repository,
	adapter Adapter,
) ServiceBuilder {
	builder := NewBuilder()
	encryptorBuilder := account_encryptors.NewBuilder()
	signerFactory := signers.NewFactory()
	return createServiceBuilder(
		encryptor,
		builder,
		repository,
		adapter,
		encryptorBuilder,
		signerFactory,
	)
}

// Adapter represents the account adapter
type Adapter interface {
	ToBytes(ins Account) ([]byte, error)
	ToInstance(bytes []byte) (Account, error)
}

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithEncryptor(encryptor account_encryptors.Encryptor) Builder
	WithSigner(signer signers.Signer) Builder
	Now() (Account, error)
}

// Account represents the identity account
type Account interface {
	Username() string
	Encryptor() account_encryptors.Encryptor
	Signer() signers.Signer
}

// UpdateCriteriaBuilder represents an update criteria builder
type UpdateCriteriaBuilder interface {
	Create() UpdateCriteriaBuilder
	WithUsername(username string) UpdateCriteriaBuilder
	WithPassword(password []byte) UpdateCriteriaBuilder
	ChangeSigner() UpdateCriteriaBuilder
	ChangeEncryptor() UpdateCriteriaBuilder
	Now() (UpdateCriteria, error)
}

// UpdateCriteria represents an update criteria
type UpdateCriteria interface {
	ChangeSigner() bool
	ChangeEncryptor() bool
	HasUsername() bool
	Username() string
	HasPassword() bool
	Password() []byte
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithDatabase(db databases.Database) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents the account repository
type Repository interface {
	List() ([]string, error)
	Exists(username string) (bool, error)
	Retrieve(credentials credentials.Credentials) (Account, error)
}

// ServiceBuilder represents the service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithDatabase(db databases.Database) ServiceBuilder
	WithBitrate(bitrate int) ServiceBuilder
	Now() (Service, error)
}

// Service represents the account service
type Service interface {
	Insert(account Account, password []byte) error
	Update(credentials credentials.Credentials, criteria UpdateCriteria) error
	Delete(credentials credentials.Credentials) error
}
