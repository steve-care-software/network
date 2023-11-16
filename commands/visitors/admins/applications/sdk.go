package applications

import "steve.care/8web/resources/credentials"

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the admin application
type Application interface {
}
