package applications

import (
	"errors"

	"steve.care/network/applications/applications"
	database_applications "steve.care/network/applications/databases"
	"steve.care/network/domain/databases"
)

type application struct {
	appBuilder      applications.Builder
	dbApp           database_applications.Application
	bitrate         int
	currentDbApp    databases.Database
	currentTrxApp   databases.Transaction
	currentQueryApp databases.Query
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

	database, err := app.dbApp.Open(name)
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
	if app.isActive() {
		return nil, errors.New(currentActiveErrorMsg)
	}

	database, err := app.dbApp.OpenInMemory()
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
	if app.isActive() {
		return nil, errors.New(currentActiveErrorMsg)
	}

	database, err := app.dbApp.Open(name)
	if err != nil {
		return nil, err
	}

	return app.begin(database)
}

// BeginInMemory begins the application in memory
func (app *application) BeginInMemory() (applications.Application, error) {
	if app.isActive() {
		return nil, errors.New(currentActiveErrorMsg)
	}

	database, err := app.dbApp.OpenInMemory()
	if err != nil {
		return nil, err
	}

	return app.begin(database)
}

// Commit commits the application
func (app *application) Commit() error {
	if !app.isActive() {
		return errors.New(notActiveErrorMsg)
	}

	err := app.currentTrxApp.Commit()
	if err != nil {
		return err
	}

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
	trxApp, err := database.Prepare()
	if err != nil {
		return nil, err
	}

	query := database.Query()
	appIns, err := app.appBuilder.Create().
		WithBitrate(app.bitrate).
		WithQuery(query).
		WithTransaction(trxApp).
		Now()

	if err != nil {
		return nil, err
	}

	app.currentTrxApp = trxApp
	app.currentQueryApp = query
	app.currentDbApp = database
	return appIns, nil
}

func (app *application) isActive() bool {
	return app.currentDbApp != nil &&
		app.currentQueryApp != nil &&
		app.currentTrxApp != nil
}
