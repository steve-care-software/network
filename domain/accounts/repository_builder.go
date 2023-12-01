package accounts

import (
	"errors"

	"steve.care/network/domain/databases"
	"steve.care/network/domain/encryptors"
)

type repositoryBuilder struct {
	encryptor encryptors.Encryptor
	adapter   Adapter
	db        databases.Database
}

func createRepositoryBuilder(
	encryptor encryptors.Encryptor,
	adapter Adapter,
) RepositoryBuilder {
	out := repositoryBuilder{
		encryptor: encryptor,
		adapter:   adapter,
		db:        nil,
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

// WithDatabase adds a database to the builder
func (app *repositoryBuilder) WithDatabase(db databases.Database) RepositoryBuilder {
	app.db = db
	return app
}

// Now builds a new Repository instance
func (app *repositoryBuilder) Now() (Repository, error) {
	if app.db == nil {
		return nil, errors.New("the database is mandatory in order to build a Repository instance")
	}

	return createRepository(
		app.encryptor,
		app.adapter,
		app.db,
	), nil

}
