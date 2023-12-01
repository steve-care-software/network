package accounts

import (
	"errors"

	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases"
	"steve.care/network/domain/encryptors"
)

type repository struct {
	encryptor encryptors.Encryptor
	adapter   Adapter
	db        databases.Database
}

func createRepository(
	encryptor encryptors.Encryptor,
	adapter Adapter,
	db databases.Database,
) Repository {
	out := repository{
		encryptor: encryptor,
		adapter:   adapter,
		db:        db,
	}

	return &out
}

// List returns the list of usernames
func (app *repository) List() ([]string, error) {
	uncasted, err := app.db.Query(
		func(row databases.Scannable) (interface{}, error) {
			var retUsername string
			err := row.Scan(&retUsername)
			if err != nil {
				return nil, err
			}

			return retUsername, nil
		},
		"SELECT username FROM accounts",
	)

	if err != nil {
		return nil, err
	}

	usernames := []string{}
	for _, oneIns := range uncasted {
		if str, ok := oneIns.(string); ok {
			usernames = append(usernames, str)
			continue
		}

		return nil, errors.New("the username could not be casted to a string")
	}

	return usernames, nil
}

// Exists returns true if the username exists, false otherwise
func (app *repository) Exists(username string) (bool, error) {
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
func (app *repository) Retrieve(credentials credentials.Credentials) (Account, error) {
	username := credentials.Username()
	uncasted, err := app.db.QueryFirst(
		func(row databases.Scannable) (interface{}, error) {
			var retCipher []byte
			err := row.Scan(&retCipher)
			if err != nil {
				return nil, err
			}

			return retCipher, nil
		},
		"SELECT cipher FROM accounts WHERE username = ?",
		username,
	)

	if err != nil {
		return nil, err
	}

	if cipher, ok := uncasted.([]byte); ok {
		password := credentials.Password()
		retBytes, err := app.encryptor.Decrypt(cipher, password)
		if err != nil {
			return nil, err
		}

		return app.adapter.ToInstance(retBytes)
	}

	return nil, errors.New("the cipher could not be casted to []byte")

}
