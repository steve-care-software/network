package sqllites

import (
	"database/sql"
	"errors"
	"fmt"

	"steve.care/network/domain/accounts"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/encryptors"
)

type accountRepository struct {
	encryptor encryptors.Encryptor
	adapter   accounts.Adapter
	dbPtr     *sql.DB
}

func createAccountRepository(
	encryptor encryptors.Encryptor,
	adapter accounts.Adapter,
	dbPtr *sql.DB,
) accounts.Repository {
	out := accountRepository{
		encryptor: encryptor,
		adapter:   adapter,
		dbPtr:     dbPtr,
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
		err := rows.Scan(&retUsername)
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
	rows, err := app.dbPtr.Query("SELECT cipher FROM accounts WHERE username = ?", username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		str := fmt.Sprintf("the given username (%s) do NOT match an account", username)
		return nil, errors.New(str)
	}

	var retCipher []byte
	err = rows.Scan(&retCipher)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	password := credentials.Password()
	retBytes, err := app.encryptor.Decrypt(retCipher, password)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToInstance(retBytes)
}
