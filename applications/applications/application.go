package applications

import (
	accounts_application "steve.care/network/applications/applications/accounts"
	resources_application "steve.care/network/applications/applications/resources"
)

type application struct {
	accountApplication  accounts_application.Application
	resourceApplication resources_application.Application
}

func createApplication(
	accountApplication accounts_application.Application,
	resourceApplication resources_application.Application,
) Application {
	out := application{
		accountApplication:  accountApplication,
		resourceApplication: resourceApplication,
	}

	return &out
}

// Accounts returns the account application
func (app *application) Accounts() accounts_application.Application {
	return app.accountApplication
}

// Resources returns the resource application
func (app *application) Resources() resources_application.Application {
	return app.resourceApplication
}
