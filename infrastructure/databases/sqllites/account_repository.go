package sqllites

import (
	"database/sql"

	"steve.care/network/applications/applications/encryptors"
	"steve.care/network/domain/accounts"
	"steve.care/network/domain/credentials"
)

type accountRepository struct {
	encryptorApp encryptors.Application
	adapter      accounts.Adapter
	dbPtr        *sql.DB
}

func createAccountRepository(
	encryptorApp encryptors.Application,
	adapter accounts.Adapter,
	dbPtr *sql.DB,
) accounts.Repository {
	out := accountRepository{
		encryptorApp: encryptorApp,
		adapter:      adapter,
		dbPtr:        dbPtr,
	}

	return &out
}

// List returns the list of usernames
func (app *accountRepository) List() ([]string, error) {
	rows, err := app.dbPtr.Query("SELECT username FROM accounts")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	usernames := []string{}
	for rows.Next() {
		var retUsername string
		err = rows.Scan(&retUsername)
		if err != nil {
			return nil, err
		}

		usernames = append(usernames, retUsername)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return usernames, nil
}

// Exists returns true if the username exists, false otherwise
func (app *accountRepository) Exists(username string) (bool, error) {
	names, err := app.List()
	if err != nil {
		return false, err
	}

	for _, oneName := range names {
		if oneName != username {
			continue
		}

		return true, nil
	}

	return false, nil
}

// Retrieve retrieves the accont from credentials
func (app *accountRepository) Retrieve(credentials credentials.Credentials) (accounts.Account, error) {
	username := credentials.Username()
	row := app.dbPtr.QueryRow("SELECT cipher FROM accounts WHERE username = ?", username)

	var retCipher []byte
	err := row.Scan(&retCipher)
	if err != nil {
		return nil, err
	}

	password := credentials.Password()
	retBytes, err := app.encryptorApp.Decrypt(retCipher, password)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToInstance(retBytes)
}
