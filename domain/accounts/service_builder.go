package accounts

import (
	"errors"

	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/databases"
	"steve.care/network/domain/encryptors"
)

type serviceBuilder struct {
	encryptor        encryptors.Encryptor
	builder          Builder
	repository       Repository
	adapter          Adapter
	encryptorBuilder account_encryptors.Builder
	signerFactory    signers.Factory
	db               databases.Database
	bitrate          int
}

func createServiceBuilder(
	encryptor encryptors.Encryptor,
	builder Builder,
	repository Repository,
	adapter Adapter,
	encryptorBuilder account_encryptors.Builder,
	signerFactory signers.Factory,
) ServiceBuilder {
	out := serviceBuilder{
		encryptor:        encryptor,
		builder:          builder,
		repository:       repository,
		adapter:          adapter,
		encryptorBuilder: encryptorBuilder,
		signerFactory:    signerFactory,
		db:               nil,
		bitrate:          0,
	}

	return &out
}

// Create initializes the builder
func (app *serviceBuilder) Create() ServiceBuilder {
	return createServiceBuilder(
		app.encryptor,
		app.builder,
		app.repository,
		app.adapter,
		app.encryptorBuilder,
		app.signerFactory,
	)
}

// WithDatabase adds a database to the builder
func (app *serviceBuilder) WithDatabase(db databases.Database) ServiceBuilder {
	app.db = db
	return app
}

// WithBitrate adds a bitrate to the builder
func (app *serviceBuilder) WithBitrate(bitrate int) ServiceBuilder {
	app.bitrate = bitrate
	return app
}

// Now builds a new Service instance
func (app *serviceBuilder) Now() (Service, error) {
	if app.db == nil {
		return nil, errors.New("the database is mandatory in order to build a Service instance")
	}

	if app.bitrate <= 0 {
		return nil, errors.New("the bitrate must be greater than zero (0) in order to build a Service instance")
	}

	return createService(
		app.encryptor,
		app.builder,
		app.repository,
		app.adapter,
		app.encryptorBuilder,
		app.signerFactory,
		app.db,
		app.bitrate,
	), nil

}
