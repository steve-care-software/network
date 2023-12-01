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
	trx               databases.Transaction
	query             databases.Query
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
		trx:               nil,
		query:             nil,
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

// WithQuery adds a query to the builder
func (app *builder) WithQuery(query databases.Query) Builder {
	app.query = query
	return app
}

// WithTransaction adds a trx to the builder
func (app *builder) WithTransaction(trx databases.Transaction) Builder {
	app.trx = trx
	return app
}

// WithBitrate adds a bitrate to the builder
func (app *builder) WithBitrate(bitrate int) Builder {
	app.bitrate = bitrate
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.query == nil {
		return nil, errors.New("the query is mandatory in order to build an Application instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transaction is mandatory in order to build an Application instance")
	}

	if app.bitrate <= 0 {
		return nil, errors.New("the bitrate is mandatory in order to build an Application instance")
	}

	repository, err := app.repositoryBuilder.Create().
		WithQuery(app.query).
		Now()

	if err != nil {
		return nil, err
	}

	service, err := app.serviceBuilder.Create().
		WithQuery(app.query).
		WithTransaction(app.trx).
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
