package sqllites

import (
	"steve.care/network/applications"
	applications_applications "steve.care/network/applications/applications"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/encryptors"
)

const notActiveErrorMsg = "the application NEVER began a transactional state, therefore that method cannot be executed"
const currentActiveErrorMsg = "the application ALREADY began a transactional state, therefore that method cannot be executed"

// NewApplication creates a new application
func NewApplication(
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	bitrate int,
	basePath string,
) applications.Application {
	appBuilder := applications_applications.NewBuilder(
		encryptor,
		adapter,
	)

	return createApplication(
		appBuilder,
		bitrate,
		basePath,
	)
}
