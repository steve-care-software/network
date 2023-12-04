package applications

import (
	"steve.care/network/applications/applications/accounts"
	accounts_application "steve.care/network/applications/applications/accounts"
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/databases/transactions"
)

type application struct {
	accAppBuilder accounts.Builder
	trx           transactions.Transaction
	query         queries.Query
	bitrate       int
}

func createApplication(
	accAppBuilder accounts.Builder,
	trx transactions.Transaction,
	query queries.Query,
	bitrate int,
) Application {
	out := application{
		accAppBuilder: accAppBuilder,
		trx:           trx,
		query:         query,
		bitrate:       bitrate,
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
