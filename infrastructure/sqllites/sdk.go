package sqllites

import (
	"database/sql"

	"steve.care/network/applications"
	"steve.care/network/domain/accounts"
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/encryptors"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	token_dashboards "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/dashboards"
	commands_layers "steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/schemas"
	"steve.care/network/domain/schemas/groups"
	schema_resources "steve.care/network/domain/schemas/groups/resources"
	"steve.care/network/domain/schemas/groups/resources/fields"
	field_methods "steve.care/network/domain/schemas/groups/resources/fields/methods"
	field_types "steve.care/network/domain/schemas/groups/resources/fields/types"
	"steve.care/network/domain/schemas/groups/resources/fields/types/dependencies"
	"steve.care/network/domain/schemas/groups/resources/methods"
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
	dashboardBuilder := token_dashboards.NewBuilder()
	viewportBuilder := viewports.NewBuilder()
	cmdLayerBuilder := commands_layers.NewLayerBuilder()
	builders := map[string]interface{}{
		"resources_dashboards_viewport": viewports.NewBuilder(),
		"resources_dashboards":          token_dashboards.NewBuilder(),
		"resources":                     tokens.NewContentBuilder(),
	}

	return createResourceRepository(
		hashAdapter,
		signatureAdapter,
		builder,
		tokenBuilder,
		dashboardBuilder,
		viewportBuilder,
		cmdLayerBuilder,
		builders,
		schema,
		dbPtr,
	)
}

// NewResourceService creates a new resoruce service
func NewResourceService(
	schema schemas.Schema,
	txPtr *sql.Tx,
) resources.Service {
	hashAdapter := hash.NewAdapter()
	return createResourceService(
		hashAdapter,
		schema,
		txPtr,
	)
}

// NewSchemaFactory creates a new schema factory
func NewSchemaFactory(
	keyFieldName string,
	keyFieldMethodNames []string,
) schemas.Factory {
	builder := schemas.NewBuilder()
	groupBuilder := groups.NewBuilder()
	methodChainsBuilder := groups.NewMethodChainsBuilder()
	methodChainBuilder := groups.NewMethodChainBuilder()
	elementBuilder := groups.NewElementBuilder()
	resourceBuilder := schema_resources.NewBuilder()
	resourceMethodsBuilder := methods.NewBuilder()
	connectionsBuilder := schema_resources.NewConnectionsBuilder()
	connectionBuilder := schema_resources.NewConnectionBuilder()
	pointerBuilder := schema_resources.NewPointerBuilder()
	fieldsBuilder := fields.NewBuilder()
	fieldBuilder := fields.NewFieldBuilder()
	fieldMethodsBuilder := field_methods.NewBuilder()
	fieldTypeBuilder := field_types.NewBuilder()
	fieldDependencyBuilder := dependencies.NewBuilder()
	return createSchemaFactory(
		builder,
		groupBuilder,
		methodChainsBuilder,
		methodChainBuilder,
		elementBuilder,
		resourceBuilder,
		resourceMethodsBuilder,
		connectionsBuilder,
		connectionBuilder,
		pointerBuilder,
		fieldsBuilder,
		fieldBuilder,
		fieldMethodsBuilder,
		fieldTypeBuilder,
		fieldDependencyBuilder,
		keyFieldName,
		keyFieldMethodNames,
	)
}
