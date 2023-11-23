package applications

import (
	"steve.care/network/applications/authenticates"
	"steve.care/network/domain/credentials"
)

// Application represents a stencil application
type Application interface {
	List() ([]string, error)
	Insert(credentials credentials.Credentials) error
	Authenticate(credentials credentials.Credentials) (authenticates.Application, error)
}
