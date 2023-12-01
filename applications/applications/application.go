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
	database       databases.Database
	bitrate        int
}

func createApplication(
	accAppBuilder accounts.Builder,
	authAppBuilder authenticates.Builder,
	database databases.Database,
	bitrate int,
) Application {
	out := application{
		accAppBuilder:  accAppBuilder,
		authAppBuilder: authAppBuilder,
		database:       database,
		bitrate:        bitrate,
	}

	return &out
}

// Accounts returns the account application
func (app *application) Accounts() (accounts_application.Application, error) {
	return app.accAppBuilder.Create().
		WithBitrate(app.bitrate).
		WithDatabase(app.database).
		Now()
}

// Authenticate returns the authenticate application
func (app *application) Authenticate(credentials credentials.Credentials) (authenticates.Application, error) {
	return app.authAppBuilder.Create().
		WithCredentials(credentials).
		WithDatabase(app.database).
		Now()
}
