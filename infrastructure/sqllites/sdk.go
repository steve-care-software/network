package sqllites

import (
	"database/sql"

	"steve.care/network/applications"
	"steve.care/network/domain/accounts"
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/dashboards/widgets"
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/encryptors"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/orms"
	"steve.care/network/domain/orms/schemas"
	"steve.care/network/domain/orms/schemas/roots"
	"steve.care/network/domain/orms/schemas/roots/groups"
	"steve.care/network/domain/orms/schemas/roots/groups/methods"
	schema_resources "steve.care/network/domain/orms/schemas/roots/groups/resources"
	"steve.care/network/domain/orms/schemas/roots/groups/resources/fields"
	field_methods "steve.care/network/domain/orms/schemas/roots/groups/resources/fields/methods"
	field_types "steve.care/network/domain/orms/schemas/roots/groups/resources/fields/types"
	"steve.care/network/domain/orms/schemas/roots/groups/resources/fields/types/dependencies"
	root_methods "steve.care/network/domain/orms/schemas/roots/methods"
	"steve.care/network/domain/orms/skeletons"
	"steve.care/network/domain/orms/skeletons/connections"
	skeleton_resources "steve.care/network/domain/orms/skeletons/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	token_dashboards "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/dashboards"
	commands_layers "steve.care/network/domain/programs/logics/libraries/layers"
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

// NewOrmRepository creates a new orm repository
func NewOrmRepository(
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
) orms.Repository {
	hashAdapter := hash.NewAdapter()
	builders := map[string]interface{}{
		"tokens_dashboards_widget":   widgets.NewWidgetBuilder(),
		"tokens_dashboards_viewport": viewports.NewBuilder(),
		"tokens_dashboards":          token_dashboards.NewBuilder(),
		"tokens":                     tokens.NewContentBuilder(),
	}

	return createOrmRepository(
		hashAdapter,
		builders,
		skeleton,
		dbPtr,
	)
}

// NewOrmService creates a new orm service
func NewOrmService(
	skeleton skeletons.Skeleton,
	txPtr *sql.Tx,
) orms.Service {
	hashAdapter := hash.NewAdapter()
	return createOrmService(
		hashAdapter,
		skeleton,
		txPtr,
	)
}

// NewSkeletonFactory creates a new skeleton factory
func NewSkeletonFactory() skeletons.Factory {
	builder := skeletons.NewBuilder()
	resourcesBuilder := skeleton_resources.NewBuilder()
	resourceBuilder := skeleton_resources.NewResourceBuilder()
	fieldsBuilder := skeleton_resources.NewFieldsBuilder()
	fieldBuilder := skeleton_resources.NewFieldBuilder()
	kindBuilder := skeleton_resources.NewKindBuilder()
	connectionsBuilder := connections.NewBuilder()
	connectionBuilder := connections.NewConnectionBuilder()
	return createSkeletonFactory(
		builder,
		resourcesBuilder,
		resourceBuilder,
		fieldsBuilder,
		fieldBuilder,
		kindBuilder,
		connectionsBuilder,
		connectionBuilder,
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
		"tokens_dashboards_widget":   widgets.NewWidgetBuilder(),
		"tokens_dashboards_viewport": viewports.NewBuilder(),
		"tokens_dashboards":          token_dashboards.NewBuilder(),
		"tokens":                     tokens.NewContentBuilder(),
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
	rootsBuilder := roots.NewBuilder()
	rootBuilder := roots.NewRootBuilder()
	rootMethodBuilder := root_methods.NewBuilder()
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
		rootsBuilder,
		rootBuilder,
		rootMethodBuilder,
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
