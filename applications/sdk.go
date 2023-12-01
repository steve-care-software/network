package applications

import (
	"steve.care/network/applications/applications"
	"steve.care/network/applications/databases"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/encryptors"
)

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
	InitInMemory(name string, script string) (applications.Application, error)
	Begin(name string) (applications.Application, error)
	BeginInMemory(name string) (applications.Application, error)
	Commit(name string) error
	Cancel(name string) error
	Rollback(name string) error
	Close(name string) error
}
