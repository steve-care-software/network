package applications

import (
	"steve.care/network/applications/applications/accounts"
	accounts_application "steve.care/network/applications/applications/accounts"
	"steve.care/network/applications/applications/authenticates"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases"
)

type application struct {
	accAppBuilder  accounts.Builder
	authAppBuilder authenticates.Builder
	trx            databases.Transaction
	query          databases.Query
	bitrate        int
}

func createApplication(
	accAppBuilder accounts.Builder,
	authAppBuilder authenticates.Builder,
	trx databases.Transaction,
	query databases.Query,
	bitrate int,
) Application {
	out := application{
		accAppBuilder:  accAppBuilder,
		authAppBuilder: authAppBuilder,
		trx:            trx,
		query:          query,
		bitrate:        bitrate,
	}

	return &out
}

// Accounts returns the account application
func (app *application) Accounts() (accounts_application.Application, error) {
	return app.accAppBuilder.Create().
		WithBitrate(app.bitrate).
		WithQuery(app.query).
		WithTransaction(app.trx).
		Now()
}

// Authenticate returns the authenticate application
func (app *application) Authenticate(credentials credentials.Credentials) (authenticates.Application, error) {
	return app.authAppBuilder.Create().
		WithCredentials(credentials).
		WithTransaction(app.trx).
		WithQuery(app.query).
		Now()
}
