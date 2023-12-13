package applications

import (
	accounts_application "steve.care/network/applications/applications/accounts"
)

type application struct {
	accountApplication accounts_application.Application
}

func createApplication(
	accountApplication accounts_application.Application,
) Application {
	out := application{
		accountApplication: accountApplication,
	}

	return &out
}

// Accounts returns the account application
func (app *application) Accounts() accounts_application.Application {
	return app.accountApplication
}
