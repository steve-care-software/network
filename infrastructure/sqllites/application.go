package sqllites

import (
	"database/sql"
	"path/filepath"

	db_applications "steve.care/network/applications/databases"
	"steve.care/network/domain/databases"
)

type application struct {
	basePath string
}

func createApplication(
	basePath string,
) db_applications.Application {
	out := application{
		basePath: basePath,
	}

	return &out
}

// OpenInMemory opens a database in memory
func (app *application) OpenInMemory() (databases.Database, error) {
	basePath := filepath.Join(app.basePath, ":memory:")
	dbPtr, err := sql.Open("sqlite3", basePath)
	if err != nil {
		return nil, err
	}

	return createDatabase(
		createQuery(dbPtr),
		dbPtr,
	), nil
}

// Open opens a database
func (app *application) Open(name string) (databases.Database, error) {
	basePath := filepath.Join(app.basePath, name)
	dbPtr, err := sql.Open("sqlite3", basePath)
	if err != nil {
		return nil, err
	}

	return createDatabase(
		createQuery(dbPtr),
		dbPtr,
	), nil
}
