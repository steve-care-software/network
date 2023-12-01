package accounts

import (
	"errors"

	"steve.care/network/domain/accounts"
	"steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/databases"
)

type builder struct {
	accountBuilder    accounts.Builder
	signerFactory     signers.Factory
	encryptorBuilder  encryptors.Builder
	repositoryBuilder accounts.RepositoryBuilder
	serviceBuilder    accounts.ServiceBuilder
	database          databases.Database
	bitrate           int
}

func createBuilder(
	accountBuilder accounts.Builder,
	signerFactory signers.Factory,
	encryptorBuilder encryptors.Builder,
	repositoryBuilder accounts.RepositoryBuilder,
	serviceBuilder accounts.ServiceBuilder,
) Builder {
	out := builder{
		accountBuilder:    accountBuilder,
		signerFactory:     signerFactory,
		encryptorBuilder:  encryptorBuilder,
		repositoryBuilder: repositoryBuilder,
		serviceBuilder:    serviceBuilder,
		database:          nil,
		bitrate:           0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.accountBuilder,
		app.signerFactory,
		app.encryptorBuilder,
		app.repositoryBuilder,
		app.serviceBuilder,
	)
}

// WithDatabase adds a database to the builder
func (app *builder) WithDatabase(database databases.Database) Builder {
	app.database = database
	return app
}

// WithBitrate adds a bitrate to the builder
func (app *builder) WithBitrate(bitrate int) Builder {
	app.bitrate = bitrate
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.database == nil {
		return nil, errors.New("the database is mandatory in order to build an Application instance")
	}

	if app.bitrate <= 0 {
		return nil, errors.New("the bitrate is mandatory in order to build an Application instance")
	}

	repository, err := app.repositoryBuilder.Create().
		WithDatabase(app.database).
		Now()

	if err != nil {
		return nil, err
	}

	service, err := app.serviceBuilder.Create().
		WithDatabase(app.database).
		WithBitrate(app.bitrate).
		Now()

	if err != nil {
		return nil, err
	}

	return createApplication(
		app.accountBuilder,
		app.signerFactory,
		app.encryptorBuilder,
		repository,
		service,
		app.bitrate,
	), nil
}
