package sqllites

import (
	"steve.care/network/domain/databases"
)

func openThenPrepareSQLInMemoryForTests(script string) (databases.Database, error) {
	dbApp := NewApplication("./")
	dbOpenApp, err := dbApp.OpenInMemory()
	if err != nil {
		return nil, err
	}

	err = dbOpenApp.Execute(script)
	if err != nil {
		return nil, err
	}

	return dbOpenApp, nil
}
