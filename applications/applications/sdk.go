package applications

import (
	accounts_application "steve.care/network/applications/applications/accounts"
)

// NewApplication creates a new application
func NewApplication(
	accountApplication accounts_application.Application,
) Application {
	return createApplication(
		accountApplication,
	)
}

// Application represents a stencil application
type Application interface {
	Accounts() accounts_application.Application
}
