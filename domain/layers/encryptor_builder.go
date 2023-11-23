package layers

import (
	"errors"

	"steve.care/network/domain/hash"
)

type encryptorBuilder struct {
	hashAdapter hash.Adapter
	decrypt     BytesReference
	encrypt     BytesReference
	isPublicKey bool
}

func createEncryptorBuilder(
	hashAdapter hash.Adapter,
) EncryptorBuilder {
	out := encryptorBuilder{
		hashAdapter: hashAdapter,
		decrypt:     nil,
		encrypt:     nil,
		isPublicKey: false,
	}

	return &out
}

// Create initializes the builder
func (app *encryptorBuilder) Create() EncryptorBuilder {
	return createEncryptorBuilder(
		app.hashAdapter,
	)
}

// WithDecrypt adds a decrypt to the builder
func (app *encryptorBuilder) WithDecrypt(decrypt BytesReference) EncryptorBuilder {
	app.decrypt = decrypt
	return app
}

// WithEncrypt adds an encrypt to the builder
func (app *encryptorBuilder) WithEncrypt(encrypt BytesReference) EncryptorBuilder {
	app.encrypt = encrypt
	return app
}

// IsPublicKey flags the builder as isPublicKey
func (app *encryptorBuilder) IsPublicKey() EncryptorBuilder {
	app.isPublicKey = true
	return app
}

// Now builds a new Encryptor instance
func (app *encryptorBuilder) Now() (Encryptor, error) {
	data := [][]byte{}
	if app.decrypt != nil {
		data = append(data, []byte("decrypt"))
		data = append(data, app.decrypt.Hash().Bytes())
	}

	if app.encrypt != nil {
		data = append(data, []byte("encrypt"))
		data = append(data, app.encrypt.Hash().Bytes())
	}

	if app.isPublicKey {
		data = append(data, []byte("isPublicKey"))
	}

	if len(data) <= 0 {
		return nil, errors.New("the Encryptor is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.decrypt != nil {
		return createEncryptorWithDecrypt(*pHash, app.decrypt), nil
	}

	if app.encrypt != nil {
		return createEncryptorWithEncrypt(*pHash, app.encrypt), nil
	}

	return createEncryptorWithIsPublicKey(*pHash), nil
}
