package accounts

import (
	"errors"

	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/databases"
	"steve.care/network/domain/encryptors"
)

type serviceBuilder struct {
	encryptor         encryptors.Encryptor
	builder           Builder
	repositoryBuilder RepositoryBuilder
	adapter           Adapter
	encryptorBuilder  account_encryptors.Builder
	signerFactory     signers.Factory
	query             databases.Query
	trx               databases.Transaction
	bitrate           int
}

func createServiceBuilder(
	encryptor encryptors.Encryptor,
	builder Builder,
	repositoryBuilder RepositoryBuilder,
	adapter Adapter,
	encryptorBuilder account_encryptors.Builder,
	signerFactory signers.Factory,
) ServiceBuilder {
	out := serviceBuilder{
		encryptor:         encryptor,
		builder:           builder,
		repositoryBuilder: repositoryBuilder,
		adapter:           adapter,
		encryptorBuilder:  encryptorBuilder,
		signerFactory:     signerFactory,
		query:             nil,
		trx:               nil,
		bitrate:           0,
	}

	return &out
}

// Create initializes the builder
func (app *serviceBuilder) Create() ServiceBuilder {
	return createServiceBuilder(
		app.encryptor,
		app.builder,
		app.repositoryBuilder,
		app.adapter,
		app.encryptorBuilder,
		app.signerFactory,
	)
}

// WithQuery adds a query to the builder
func (app *serviceBuilder) WithQuery(query databases.Query) ServiceBuilder {
	app.query = query
	return app
}

// WithTransaction adds a trx to the builder
func (app *serviceBuilder) WithTransaction(trx databases.Transaction) ServiceBuilder {
	app.trx = trx
	return app
}

// WithBitrate adds a bitrate to the builder
func (app *serviceBuilder) WithBitrate(bitrate int) ServiceBuilder {
	app.bitrate = bitrate
	return app
}

// Now builds a new Service instance
func (app *serviceBuilder) Now() (Service, error) {
	if app.query == nil {
		return nil, errors.New("the query is mandatory in order to build a Service instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transaction is mandatory in order to build a Service instance")
	}

	if app.bitrate <= 0 {
		return nil, errors.New("the bitrate must be greater than zero (0) in order to build a Service instance")
	}

	repository, err := app.repositoryBuilder.Create().
		WithQuery(app.query).
		Now()

	if err != nil {
		return nil, err
	}

	return createService(
		app.encryptor,
		app.builder,
		repository,
		app.adapter,
		app.encryptorBuilder,
		app.signerFactory,
		app.trx,
		app.bitrate,
	), nil

}
