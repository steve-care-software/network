package applications

import (
	accounts_application "steve.care/network/applications/applications/accounts"
	resources_application "steve.care/network/applications/applications/resources"
)

// NewApplication creates a new application
func NewApplication(
	accountApplication accounts_application.Application,
	resourceApplication resources_application.Application,
) Application {
	return createApplication(
		accountApplication,
		resourceApplication,
	)
}

// Application represents a stencil application
type Application interface {
	Accounts() accounts_application.Application
	Resources() resources_application.Application
}
