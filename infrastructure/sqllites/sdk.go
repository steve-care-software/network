package sqllites

import (
	"database/sql"

	"steve.care/network/applications"
	"steve.care/network/domain/accounts"
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/encryptors"
	"steve.care/network/domain/hash"
	commands_layers "steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources/tokens"
	"steve.care/network/domain/programs/blockchains/blocks/executions/actions/resources/tokens/layers"
)

const notActiveErrorMsg = "the application NEVER began a transactional state, therefore that method cannot be executed"
const currentActiveErrorMsg = "the application ALREADY began a transactional state, therefore that method cannot be executed"

const timeLayout = "2006-01-02T15:04:05.999999999Z07:00"

// NewApplication creates a new application
func NewApplication(
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	bitrate int,
	basePath string,
) applications.Application {
	return createApplication(
		encryptor,
		adapter,
		bitrate,
		basePath,
	)
}

// NewAccountRepository creates a new account repository
func NewAccountRepository(
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	dbPtr *sql.DB,
) accounts.Repository {
	return createAccountRepository(
		encryptor,
		adapter,
		dbPtr,
	)
}

// NewAccountService creates a new account service
func NewAccountService(
	repository accounts.Repository,
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	bitrate int,
	txPtr *sql.Tx,
) accounts.Service {
	builder := accounts.NewBuilder()
	encryptorBuilder := account_encryptors.NewBuilder()
	signerFactory := signers.NewFactory()
	return createAccountService(
		encryptor,
		builder,
		repository,
		adapter,
		encryptorBuilder,
		signerFactory,
		bitrate,
		txPtr,
	)
}

// NewResourceRepository creates a new resource repository
func NewResourceRepository(
	dbPtr *sql.DB,
) resources.Repository {
	hashAdapter := hash.NewAdapter()
	signatureAdapter := signers.NewSignatureAdapter()
	builder := resources.NewBuilder()
	tokenBuilder := tokens.NewBuilder()
	layerBuilder := layers.NewBuilder()
	cmdLayerBuilder := commands_layers.NewBuilder()
	cmdLayerBytesReferenceBuilder := commands_layers.NewBytesReferenceBuilder()
	return createResourceRepository(
		hashAdapter,
		signatureAdapter,
		builder,
		tokenBuilder,
		layerBuilder,
		cmdLayerBuilder,
		cmdLayerBytesReferenceBuilder,
		dbPtr,
	)
}

// NewResourceService creates a new resoruce service
func NewResourceService(
	txPtr *sql.Tx,
) resources.Service {
	return createResourceService(
		txPtr,
	)
}
