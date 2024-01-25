package sqllites

import (
	"database/sql"

	"steve.care/network/applications"
	"steve.care/network/domain/accounts"
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/encryptors"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/layers"
	commands_layers "steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/schemas"
	"steve.care/network/domain/schemas/groups"
	schema_resources "steve.care/network/domain/schemas/groups/resources"
	"steve.care/network/domain/schemas/groups/resources/fields"
)

const notActiveErrorMsg = "the application NEVER began a transactional state, therefore that method cannot be executed"
const currentActiveErrorMsg = "the application ALREADY began a transactional state, therefore that method cannot be executed"

const timeLayout = "2006-01-02T15:04:05.999999999Z07:00"

const groupNameDelimiterForTableNames = "_"

// NewApplication creates a new application
func NewApplication(
	schema schemas.Schema,
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	bitrate int,
	basePath string,
	baseSchema string,
) applications.Application {
	return createApplication(
		schema,
		encryptor,
		adapter,
		bitrate,
		basePath,
		baseSchema,
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
	schema schemas.Schema,
	dbPtr *sql.DB,
) resources.Repository {
	hashAdapter := hash.NewAdapter()
	signatureAdapter := signers.NewSignatureAdapter()
	builder := resources.NewBuilder()
	tokenBuilder := tokens.NewBuilder()
	layerBuilder := layers.NewBuilder()
	cmdLayerBuilder := commands_layers.NewLayerBuilder()
	return createResourceRepository(
		hashAdapter,
		signatureAdapter,
		builder,
		tokenBuilder,
		layerBuilder,
		cmdLayerBuilder,
		schema,
		dbPtr,
	)
}

// NewResourceService creates a new resoruce service
func NewResourceService(
	schema schemas.Schema,
	txPtr *sql.Tx,
) resources.Service {
	return createResourceService(
		schema,
		txPtr,
	)
}

// NewSchemaFactory creates a new schema factory
func NewSchemaFactory(
	keyFieldName string,
) schemas.Factory {
	builder := schemas.NewBuilder()
	groupsBuilder := groups.NewBuilder()
	groupBuilder := groups.NewGroupBuilder()
	elementsBuilder := groups.NewElementsBuilder()
	elementBuilder := groups.NewElementBuilder()
	resourcesBuilder := schema_resources.NewBuilder()
	resourceBuilder := schema_resources.NewResourceBuilder()
	connectionsBuilder := schema_resources.NewConnectionsBuilder()
	connectionBuilder := schema_resources.NewConnectionBuilder()
	pointerBuilder := schema_resources.NewPointerBuilder()
	fieldsBuilder := fields.NewBuilder()
	fieldBuilder := fields.NewFieldBuilder()
	return createSchemaFactory(
		builder,
		groupsBuilder,
		groupBuilder,
		elementsBuilder,
		elementBuilder,
		resourcesBuilder,
		resourceBuilder,
		connectionsBuilder,
		connectionBuilder,
		pointerBuilder,
		fieldsBuilder,
		fieldBuilder,
		keyFieldName,
	)
}
