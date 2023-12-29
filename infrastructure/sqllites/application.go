package sqllites

import (
	"database/sql"
	"errors"
	"path/filepath"

	"steve.care/network/applications"
	core_applications "steve.care/network/applications/applications"
	accounts_applications "steve.care/network/applications/applications/accounts"
	resources_applications "steve.care/network/applications/applications/resources"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/encryptors"
)

type application struct {
	encryptor  encryptors.Encryptor
	adapter    accounts.Adapter
	bitrate    int
	basePath   string
	currentDb  *sql.DB
	currentTrx *sql.Tx
}

func createApplication(
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	bitrate int,
	basePath string,
) applications.Application {
	out := application{
		encryptor:  encryptor,
		adapter:    adapter,
		bitrate:    bitrate,
		basePath:   basePath,
		currentDb:  nil,
		currentTrx: nil,
	}

	return &out
}

// Init inits an application with a script
func (app *application) Init(name string, script string) (core_applications.Application, error) {
	if app.isActive() {
		return nil, errors.New(currentActiveErrorMsg)
	}

	err := app.openDatabaseIfNotAlready(name)
	if err != nil {
		return nil, err
	}

	_, err = app.currentDb.Exec(script)
	if err != nil {
		return nil, err
	}

	return app.begin()
}

// Begin begins the application
func (app *application) Begin(name string) (core_applications.Application, error) {
	err := app.openDatabaseIfNotAlready(name)
	if err != nil {
		return nil, err
	}

	return app.begin()
}

// Commit commits the application
func (app *application) Commit() error {
	err := app.currentTrx.Commit()
	if err != nil {
		return err
	}

	app.currentTrx = nil
	return nil
}

// Rollback rollbacks the application
func (app *application) Rollback() error {
	if !app.isActive() {
		return errors.New(notActiveErrorMsg)
	}

	err := app.currentTrx.Rollback()
	if err != nil {
		return err
	}

	app.currentTrx = nil
	return nil
}

// Close closes the application
func (app *application) Close() error {
	if !app.isActive() {
		return errors.New(notActiveErrorMsg)
	}

	err := app.currentDb.Close()
	if err != nil {
		return err
	}

	app.currentDb = nil
	app.currentTrx = nil
	return nil
}

func (app *application) begin() (core_applications.Application, error) {
	if !app.isTransactionActive() {
		trxApp, err := app.currentDb.Begin()
		if err != nil {
			return nil, err
		}

		app.currentTrx = trxApp
	}

	accountRepository := NewAccountRepository(
		app.encryptor,
		app.adapter,
		app.currentDb,
	)

	accountService := NewAccountService(
		accountRepository,
		app.encryptor,
		app.adapter,
		app.bitrate,
		app.currentTrx,
	)

	accountApplication := accounts_applications.NewApplication(
		accountRepository,
		accountService,
		app.bitrate,
	)

	resourceRepository := NewResourceRepository(
		app.currentDb,
	)

	resourceService := NewResourceService(
		app.currentTrx,
	)

	resourceApplication := resources_applications.NewApplication(
		resourceRepository,
		resourceService,
	)

	return core_applications.NewApplication(
		accountApplication,
		resourceApplication,
	), nil
}

func (app *application) openDatabaseIfNotAlready(name string) error {
	if !app.isDatabaseOpen() {
		database, err := app.open(name)
		if err != nil {
			return err
		}

		app.currentDb = database
	}

	return nil
}

func (app *application) isActive() bool {
	return app.isDatabaseOpen() &&
		app.currentDb != nil
}

func (app *application) isTransactionActive() bool {
	return app.currentTrx != nil
}

func (app *application) isDatabaseOpen() bool {
	return app.currentDb != nil
}

func (app *application) open(name string) (*sql.DB, error) {
	basePath := filepath.Join(app.basePath, name)
	return sql.Open("sqlite3", basePath)
}
