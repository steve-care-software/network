package sqllites

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/layers"
	commands_layers "steve.care/network/domain/programs/logics/libraries/layers"
)

type resourceRepository struct {
	hashAdapter                   hash.Adapter
	signatureAdapter              signers.SignatureAdapter
	builder                       resources.Builder
	tokenBuilder                  tokens.Builder
	layerBuilder                  layers.Builder
	cmdLayerBuilder               commands_layers.Builder
	cmdLayerBytesReferenceBuilder commands_layers.BytesReferenceBuilder
	dbPtr                         *sql.DB
}

func createResourceRepository(
	hashAdapter hash.Adapter,
	signatureAdapter signers.SignatureAdapter,
	builder resources.Builder,
	tokenBuilder tokens.Builder,
	layerBuilder layers.Builder,
	cmdLayerBuilder commands_layers.Builder,
	cmdLayerBytesReferenceBuilder commands_layers.BytesReferenceBuilder,
	dbPtr *sql.DB,
) resources.Repository {
	out := resourceRepository{
		hashAdapter:                   hashAdapter,
		signatureAdapter:              signatureAdapter,
		builder:                       builder,
		tokenBuilder:                  tokenBuilder,
		layerBuilder:                  layerBuilder,
		cmdLayerBuilder:               cmdLayerBuilder,
		cmdLayerBytesReferenceBuilder: cmdLayerBytesReferenceBuilder,
		dbPtr:                         dbPtr,
	}

	return &out
}

// Amount returns the amount of resources
func (app *resourceRepository) Amount() (uint, error) {
	return 0, nil
}

// AmountByQuery returns the amount of resources by criteria
func (app *resourceRepository) AmountByQuery(criteria hash.Hash) (uint, error) {
	return 0, nil
}

// ListByQuery lists resource hashes by criteria
func (app *resourceRepository) ListByQuery(criteria hash.Hash) ([]hash.Hash, error) {
	return nil, nil
}

// RetrieveByQuery retrieves a resource by criteria
func (app *resourceRepository) RetrieveByQuery(criteria hash.Hash) (resources.Resource, error) {
	return nil, nil
}

// RetrieveByHash retrieves a resource by hash
func (app *resourceRepository) RetrieveByHash(hash hash.Hash) (resources.Resource, error) {
	rows, err := app.dbPtr.Query("SELECT token, signature FROM resource WHERE hash = ?", hash.Bytes())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		str := fmt.Sprintf("the given hash (%s) do NOT match a Layer instance", hash.String())
		return nil, errors.New(str)
	}

	var retSignatureBytes []byte
	var retTokenHashBytes []byte
	err = rows.Scan(&retTokenHashBytes, &retSignatureBytes)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	pTokenHash, err := app.hashAdapter.FromBytes(retTokenHashBytes)
	if err != nil {
		return nil, err
	}

	token, err := app.retrieveTokenByHash(*pTokenHash)
	if err != nil {
		return nil, err
	}

	signature, err := app.signatureAdapter.ToSignature(retSignatureBytes)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithToken(token).
		WithSignature(signature).
		Now()
}

func (app *resourceRepository) retrieveTokenByHash(hash hash.Hash) (tokens.Token, error) {
	rows, err := app.dbPtr.Query("SELECT layer, created_on FROM token WHERE hash = ?", hash.Bytes())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		str := fmt.Sprintf("the given hash (%s) do NOT match a Layer instance", hash.String())
		return nil, errors.New(str)
	}

	var retCreatedOn string
	var retLayerHashBytes []byte
	err = rows.Scan(&retLayerHashBytes, &retCreatedOn)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	pLayerHash, err := app.hashAdapter.FromBytes(retLayerHashBytes)
	if err != nil {
		return nil, err
	}

	layer, err := app.retrieveLayerByHash(*pLayerHash)
	if err != nil {
		return nil, err
	}

	tm, err := time.Parse(timeLayout, retCreatedOn)
	if err != nil {
		return nil, err
	}

	return app.tokenBuilder.Create().
		WithLayer(layer).
		CreatedOn(tm).
		Now()
}

func (app *resourceRepository) retrieveLayerByHash(hash hash.Hash) (layers.Layer, error) {
	rows, err := app.dbPtr.Query("SELECT bytes_reference FROM layer WHERE hash = ?", hash.Bytes())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		str := fmt.Sprintf("the given hash (%s) do NOT match a Layer instance", hash.String())
		return nil, errors.New(str)
	}

	var retBytesReferenceHashBytes []byte
	err = rows.Scan(&retBytesReferenceHashBytes)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	builder := app.layerBuilder.Create()
	if retBytesReferenceHashBytes != nil {
		pBytesReferenceHash, err := app.hashAdapter.FromBytes(retBytesReferenceHashBytes)
		if err != nil {
			return nil, err
		}

		bytesReference, err := app.retrieveLayerBytesReferenceByHash(*pBytesReferenceHash)
		if err != nil {
			return nil, err
		}

		builder.WithBytesReference(bytesReference)
	}

	return builder.Now()
}

func (app *resourceRepository) retrieveLayerBytesReferenceByHash(hash hash.Hash) (commands_layers.BytesReference, error) {
	rows, err := app.dbPtr.Query("SELECT variable, bytes FROM layer_bytes_reference WHERE hash = ?", hash.Bytes())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		str := fmt.Sprintf("the given hash (%s) do NOT match a BytesReference instance", hash.String())
		return nil, errors.New(str)
	}

	retVariable := ""
	var retBytes []byte
	err = rows.Scan(&retVariable, &retBytes)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	builder := app.cmdLayerBytesReferenceBuilder.Create()
	if retVariable != "" {
		builder.WithVariable(retVariable)
	}

	if retBytes != nil {
		builder.WithBytes(retBytes)
	}

	return builder.Now()
}
