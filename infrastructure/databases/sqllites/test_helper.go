package sqllites

import "database/sql"

func openThenPrepareSQLInMemoryForTests(script string) (*sql.DB, error) {
	dbPtr, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	_, err = dbPtr.Exec(script)
	if err != nil {
		return nil, err
	}

	return dbPtr, nil
}
