package applications

import (
	"steve.care/network/applications/applications"
	"steve.care/network/applications/databases"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/encryptors"
)

const notActiveErrorMsg = "the application NEVER began a transactional state, therefore that method cannot be executed"
const currentActiveErrorMsg = "the application ALREADY began a transactional state, therefore that method cannot be executed"

// NewApplication creates a new application
func NewApplication(
	dbApp databases.Application,
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	bitrate int,
) Application {
	appBuilder := applications.NewBuilder(
		encryptor,
		adapter,
	)

	return createApplication(
		appBuilder,
		dbApp,
		bitrate,
	)
}

// Application represents the core application
type Application interface {
	Init(name string, script string) (applications.Application, error)
	InitInMemory(script string) (applications.Application, error)
	Begin(name string) (applications.Application, error)
	BeginInMemory() (applications.Application, error)
	Commit() error
	Cancel() error
	Rollback() error
	Close() error
}
