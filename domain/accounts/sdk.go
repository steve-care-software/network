package accounts

import (
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases/criterias/conditions"
	"steve.care/network/domain/databases/criterias/entities/resources"
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/databases/transactions"
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
	adapter Adapter,
) ServiceBuilder {
	resourceBuilder := resources.NewBuilder()
	conditionBuilder := conditions.NewBuilder()
	conditionPointerBuilder := conditions.NewPointerBuilder()
	conditionOperatorBuilder := conditions.NewOperatorBuilder()
	conditionElementBuilder := conditions.NewElementBuilder()
	conditionResourceBuilder := conditions.NewResourceBuilder()
	builder := NewBuilder()
	encryptorBuilder := account_encryptors.NewBuilder()
	signerFactory := signers.NewFactory()
	repositoryBuilder := NewRepositoryBuilder(
		encryptor,
		adapter,
	)

	return createServiceBuilder(
		resourceBuilder,
		conditionBuilder,
		conditionPointerBuilder,
		conditionOperatorBuilder,
		conditionElementBuilder,
		conditionResourceBuilder,
		encryptor,
		builder,
		repositoryBuilder,
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
	WithQuery(query queries.Query) RepositoryBuilder
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
	WithQuery(query queries.Query) ServiceBuilder
	WithTransaction(trx transactions.Transaction) ServiceBuilder
	WithBitrate(bitrate int) ServiceBuilder
	Now() (Service, error)
}

// Service represents the account service
type Service interface {
	Insert(account Account, password []byte) error
	Update(credentials credentials.Credentials, criteria UpdateCriteria) error
	Delete(credentials credentials.Credentials) error
}
