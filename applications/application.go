package applications

import (
	"errors"
	"fmt"

	"steve.care/network/applications/applications"
	database_applications "steve.care/network/applications/databases"
	"steve.care/network/domain/databases"
)

type transactionApplication struct {
	dbApp    databases.Database
	trxApp   databases.Transaction
	queryApp databases.Query
}

type application struct {
	appBuilder   applications.Builder
	dbApp        database_applications.Application
	bitrate      int
	applications map[string]transactionApplication
}

func createApplication(
	appBuilder applications.Builder,
	dbApp database_applications.Application,
	bitrate int,
) Application {
	out := application{
		appBuilder:   appBuilder,
		dbApp:        dbApp,
		bitrate:      bitrate,
		applications: map[string]transactionApplication{},
	}

	return &out
}

// Init inits an application with a script
func (app *application) Init(name string, script string) (applications.Application, error) {
	database, err := app.dbApp.Open(name)
	if err != nil {
		return nil, err
	}

	err = database.Execute(script)
	if err != nil {
		return nil, err
	}

	return app.begin(name, database)
}

// InitInMemory inits an application with a script in memory
func (app *application) InitInMemory(name string, script string) (applications.Application, error) {
	database, err := app.dbApp.OpenInMemory()
	if err != nil {
		return nil, err
	}

	err = database.Execute(script)
	if err != nil {
		return nil, err
	}

	return app.begin(name, database)
}

// Begin begins the application
func (app *application) Begin(name string) (applications.Application, error) {
	database, err := app.dbApp.Open(name)
	if err != nil {
		return nil, err
	}

	return app.begin(name, database)
}

// BeginInMemory begins the application in memory
func (app *application) BeginInMemory(name string) (applications.Application, error) {
	database, err := app.dbApp.OpenInMemory()
	if err != nil {
		return nil, err
	}

	return app.begin(name, database)
}

// Commit commits the application
func (app *application) Commit(name string) error {
	if ins, ok := app.applications[name]; ok {
		return ins.trxApp.Commit()
	}

	str := fmt.Sprintf("there is no application attached to the name: '%s', therefore the Commit operation could not execute", name)
	return errors.New(str)
}

// Cancel cancels the application
func (app *application) Cancel(name string) error {
	if ins, ok := app.applications[name]; ok {
		return ins.trxApp.Cancel()
	}

	str := fmt.Sprintf("there is no application attached to the name: '%s', therefore the Cancel operation could not execute", name)
	return errors.New(str)
}

// Rollback rollbacks the application
func (app *application) Rollback(name string) error {
	if ins, ok := app.applications[name]; ok {
		return ins.trxApp.Rollback()
	}

	str := fmt.Sprintf("there is no application attached to the name: '%s', therefore the Rollback operation could not execute", name)
	return errors.New(str)
}

// Close closes the application
func (app *application) Close(name string) error {
	if ins, ok := app.applications[name]; ok {
		err := ins.dbApp.Close()
		if err != nil {
			return err
		}

		delete(app.applications, name)
	}

	str := fmt.Sprintf("there is no application attached to the name: '%s', therefore the Close operation could not execute", name)
	return errors.New(str)
}

func (app *application) begin(name string, database databases.Database) (applications.Application, error) {
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

	app.applications[name] = transactionApplication{
		trxApp:   trxApp,
		queryApp: query,
		dbApp:    database,
	}

	return appIns, nil
}
