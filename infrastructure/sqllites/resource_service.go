package sqllites

import (
	"database/sql"

	"steve.care/network/domain/hash"
	commands_layers "steve.care/network/domain/receipts/commands/layers"
	"steve.care/network/domain/resources"
	"steve.care/network/domain/resources/tokens"
	"steve.care/network/domain/resources/tokens/layers"
	"steve.care/network/domain/resources/tokens/links"
	"steve.care/network/domain/resources/tokens/queries"
	"steve.care/network/domain/resources/tokens/receipts"
	"steve.care/network/domain/resources/tokens/suites"
)

type resourceService struct {
	txPtr *sql.Tx
}

func createResourceService(
	txPtr *sql.Tx,
) resources.Service {
	out := resourceService{
		txPtr: txPtr,
	}

	return &out
}

// Insert inserts a resource
func (app *resourceService) Insert(ins resources.Resource) error {
	token := ins.Token()
	err := app.insertToken(token)
	if err != nil {
		return err
	}

	sigBytes, err := ins.Signature().Bytes()
	if err != nil {
		return err
	}

	_, err = app.txPtr.Exec("INSERT OR IGNORE INTO resource (hash, token, signature) VALUES (?, ?, ?)", ins.Hash().Bytes(), token.Hash().Bytes(), sigBytes)
	if err != nil {
		return err
	}

	return nil
}

func (app *resourceService) insertToken(ins tokens.Token) error {
	content := ins.Content()
	if content.IsLayer() {
		layer := content.Layer()
		err := app.insertLayer(layer)
		if err != nil {
			return err
		}

		_, err = app.txPtr.Exec("INSERT OR IGNORE INTO token (hash, layer, created_on) VALUES (?, ?, ?)", ins.Hash().Bytes(), layer.Hash().Bytes(), ins.CreatedOn().Format(timeLayout))
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (app *resourceService) insertLayer(ins layers.Layer) error {
	if ins.IsBytesReference() {
		bytesReference := ins.BytesReference()
		err := app.insertLayerBytesReference(bytesReference)
		if err != nil {
			return err
		}

		_, err = app.txPtr.Exec("INSERT OR IGNORE INTO layer (hash, bytes_reference) VALUES (?, ?)", ins.Hash().Bytes(), bytesReference.Hash().Bytes())
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (app *resourceService) insertLayerBytesReference(ins commands_layers.BytesReference) error {
	hash := ins.Hash().Bytes()
	variable := ins.Variable()
	bytes := ins.Bytes()
	_, err := app.txPtr.Exec("INSERT OR IGNORE INTO layer_bytes_reference (hash, variable, bytes) VALUES (?, ?, ?)", hash, variable, bytes)
	if err != nil {
		return err
	}

	return nil
}

func (app *resourceService) insertLink(ins links.Link) error {
	return nil
}

func (app *resourceService) insertSuite(ins suites.Suite) error {
	return nil
}

func (app *resourceService) insertReceipt(ins receipts.Receipt) error {
	return nil
}

func (app *resourceService) insertQuery(ins queries.Query) error {
	return nil
}

// Delete deletes a resource
func (app *resourceService) Delete(hash hash.Hash) error {
	return nil
}
