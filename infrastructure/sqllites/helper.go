package sqllites

func getSchema() string {
	return `
		DROP TABLE IF EXISTS accounts;
		CREATE TABLE accounts(username TEXT PRIMARY KEY, cipher BLOB);
	`
}
