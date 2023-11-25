package applications

import (
	accounts_application "steve.care/network/applications/applications/accounts"
	"steve.care/network/applications/applications/authenticates"
	"steve.care/network/domain/credentials"
)

// Application represents a stencil application
type Application interface {
	Accounts() accounts_application.Application
	Authenticate(credentials credentials.Credentials) (authenticates.Application, error)
}
