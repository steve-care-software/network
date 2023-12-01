package links

import (
	"errors"

	"steve.care/network/domain/credentials"
	"steve.care/network/domain/receipts/commands/links"
)

type builder struct {
	repositoryBuilder links.RepositoryBuilder
	serviceBuilder    links.ServiceBuilder
	credentials       credentials.Credentials
}

func createBuilder(
	repositoryBuilder links.RepositoryBuilder,
	serviceBuilder links.ServiceBuilder,
) Builder {
	out := builder{
		repositoryBuilder: repositoryBuilder,
		serviceBuilder:    serviceBuilder,
		credentials:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.repositoryBuilder,
		app.serviceBuilder,
	)
}

// WithCredentials adds a credentials to the builder
func (app *builder) WithCredentials(credentials credentials.Credentials) Builder {
	app.credentials = credentials
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.credentials == nil {
		return nil, errors.New("the credentials is mandatory in order to build an Application instance")
	}

	repository, err := app.repositoryBuilder.Create().
		WithCredentials(app.credentials).
		Now()

	if err != nil {
		return nil, err
	}

	service, err := app.serviceBuilder.Create().
		WithCredentials(app.credentials).
		Now()

	if err != nil {
		return nil, err
	}

	return createApplication(
		repository,
		service,
	), nil
}
