package applications

import (
	"errors"

	"steve.care/network/applications/applications"
	database_applications "steve.care/network/applications/databases"
	"steve.care/network/domain/databases"
	"steve.care/network/domain/databases/queries"
	"steve.care/network/domain/databases/transactions"
)

type application struct {
	appBuilder      applications.Builder
	dbApp           database_applications.Application
	bitrate         int
	currentDbApp    databases.Database
	currentTrxApp   transactions.Transaction
	currentQueryApp queries.Query
}

func createApplication(
	appBuilder applications.Builder,
	dbApp database_applications.Application,
	bitrate int,
) Application {
	out := application{
		appBuilder:      appBuilder,
		dbApp:           dbApp,
		bitrate:         bitrate,
		currentDbApp:    nil,
		currentTrxApp:   nil,
		currentQueryApp: nil,
	}

	return &out
}

// Init inits an application with a script
func (app *application) Init(name string, script string) (applications.Application, error) {
	if app.isActive() {
		return nil, errors.New(currentActiveErrorMsg)
	}

	database, err := app.openDatabaseIfNotAlready(name)
	if err != nil {
		return nil, err
	}

	err = database.Execute(script)
	if err != nil {
		return nil, err
	}

	return app.begin(database)
}

// InitInMemory inits an application with a script in memory
func (app *application) InitInMemory(script string) (applications.Application, error) {
	database, err := app.openDatabaseInMemoryIfNotAlready()
	if err != nil {
		return nil, err
	}

	err = database.Execute(script)
	if err != nil {
		return nil, err
	}

	return app.begin(database)
}

// Begin begins the application
func (app *application) Begin(name string) (applications.Application, error) {
	database, err := app.openDatabaseIfNotAlready(name)
	if err != nil {
		return nil, err
	}

	return app.begin(database)
}

// BeginInMemory begins the application in memory
func (app *application) BeginInMemory() (applications.Application, error) {
	database, err := app.openDatabaseInMemoryIfNotAlready()
	if err != nil {
		return nil, err
	}

	return app.begin(database)
}

// Commit commits the application
func (app *application) Commit() error {
	err := app.currentTrxApp.Commit()
	if err != nil {
		return err
	}

	app.currentTrxApp = nil
	return nil
}

// Cancel cancels the application
func (app *application) Cancel() error {
	if !app.isActive() {
		return errors.New(notActiveErrorMsg)
	}

	err := app.currentTrxApp.Cancel()
	if err != nil {
		return err
	}

	app.currentTrxApp = nil
	return nil
}

// Rollback rollbacks the application
func (app *application) Rollback() error {
	if !app.isActive() {
		return errors.New(notActiveErrorMsg)
	}

	err := app.currentTrxApp.Rollback()
	if err != nil {
		return err
	}

	app.currentTrxApp = nil
	return nil
}

// Close closes the application
func (app *application) Close() error {
	if !app.isActive() {
		return errors.New(notActiveErrorMsg)
	}

	err := app.currentDbApp.Close()
	if err != nil {
		return err
	}

	app.currentTrxApp = nil
	app.currentQueryApp = nil
	app.currentDbApp = nil
	return nil
}

func (app *application) begin(database databases.Database) (applications.Application, error) {
	if !app.isTransactionActive() {
		trxApp, err := database.Prepare()
		if err != nil {
			return nil, err
		}

		app.currentTrxApp = trxApp
	}

	return app.appBuilder.Create().
		WithBitrate(app.bitrate).
		WithQuery(app.currentQueryApp).
		WithTransaction(app.currentTrxApp).
		Now()
}

func (app *application) openDatabaseIfNotAlready(name string) (databases.Database, error) {
	if !app.isDatabaseOpen() {
		database, err := app.dbApp.Open(name)
		if err != nil {
			return nil, err
		}

		app.currentDbApp = database
	}

	app.currentQueryApp = app.currentDbApp.Query()
	return app.currentDbApp, nil
}

func (app *application) openDatabaseInMemoryIfNotAlready() (databases.Database, error) {
	if !app.isDatabaseOpen() {
		database, err := app.dbApp.OpenInMemory()
		if err != nil {
			return nil, err
		}

		app.currentDbApp = database
	}

	app.currentQueryApp = app.currentDbApp.Query()
	return app.currentDbApp, nil
}

func (app *application) isActive() bool {
	return app.isDatabaseOpen() &&
		app.currentQueryApp != nil
}

func (app *application) isTransactionActive() bool {
	return app.currentTrxApp != nil
}

func (app *application) isDatabaseOpen() bool {
	return app.currentDbApp != nil
}
