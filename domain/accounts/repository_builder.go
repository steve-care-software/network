package accounts

import (
	"errors"

	"steve.care/network/domain/databases"
	"steve.care/network/domain/encryptors"
)

type repositoryBuilder struct {
	encryptor encryptors.Encryptor
	adapter   Adapter
	query     databases.Query
}

func createRepositoryBuilder(
	encryptor encryptors.Encryptor,
	adapter Adapter,
) RepositoryBuilder {
	out := repositoryBuilder{
		encryptor: encryptor,
		adapter:   adapter,
		query:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *repositoryBuilder) Create() RepositoryBuilder {
	return createRepositoryBuilder(
		app.encryptor,
		app.adapter,
	)
}

// WithQuery adds a query to the builder
func (app *repositoryBuilder) WithQuery(query databases.Query) RepositoryBuilder {
	app.query = query
	return app
}

// Now builds a new Repository instance
func (app *repositoryBuilder) Now() (Repository, error) {
	if app.query == nil {
		return nil, errors.New("the query is mandatory in order to build a Repository instance")
	}

	return createRepository(
		app.encryptor,
		app.adapter,
		app.query,
	), nil

}
