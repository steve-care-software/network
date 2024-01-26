package sqllites

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	token_dashboards "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/dashboards"
	commands_layers "steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/schemas"
)

type resourceRepository struct {
	hashAdapter      hash.Adapter
	signatureAdapter signers.SignatureAdapter
	builder          resources.Builder
	tokenBuilder     tokens.Builder
	dashboardBuilder token_dashboards.Builder
	viewportBuilder  viewports.Builder
	cmdLayerBuilder  commands_layers.LayerBuilder
	schema           schemas.Schema
	dbPtr            *sql.DB
}

func createResourceRepository(
	hashAdapter hash.Adapter,
	signatureAdapter signers.SignatureAdapter,
	builder resources.Builder,
	tokenBuilder tokens.Builder,
	dashboardBuilder token_dashboards.Builder,
	viewportBuilder viewports.Builder,
	cmdLayerBuilder commands_layers.LayerBuilder,
	schema schemas.Schema,
	dbPtr *sql.DB,
) resources.Repository {
	out := resourceRepository{
		hashAdapter:      hashAdapter,
		signatureAdapter: signatureAdapter,
		builder:          builder,
		tokenBuilder:     tokenBuilder,
		dashboardBuilder: dashboardBuilder,
		viewportBuilder:  viewportBuilder,
		cmdLayerBuilder:  cmdLayerBuilder,
		schema:           schema,
		dbPtr:            dbPtr,
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
	rows, err := app.dbPtr.Query("SELECT dashboards_viewport, created_on FROM token WHERE hash = ?", hash.Bytes())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		str := fmt.Sprintf("the given hash (%s) do NOT match a Layer instance", hash.String())
		return nil, errors.New(str)
	}

	var retCreatedOn string
	var retDashboardViewport []byte
	err = rows.Scan(&retDashboardViewport, &retCreatedOn)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	pDashboardViewportHash, err := app.hashAdapter.FromBytes(retDashboardViewport)
	if err != nil {
		return nil, err
	}

	dashboardViewport, err := app.retrieveDashboardViewportByHash(*pDashboardViewportHash)
	if err != nil {
		return nil, err
	}

	dashboard, err := app.dashboardBuilder.Create().WithViewport(dashboardViewport).Now()
	if err != nil {
		return nil, err
	}

	tm, err := time.Parse(timeLayout, retCreatedOn)
	if err != nil {
		return nil, err
	}

	return app.tokenBuilder.Create().
		WithDashboard(dashboard).
		CreatedOn(tm).
		Now()
}

func (app *resourceRepository) retrieveDashboardViewportByHash(hash hash.Hash) (viewports.Viewport, error) {
	rows, err := app.dbPtr.Query("SELECT row, height FROM resources_dashboards_viewport WHERE hash = ?", hash.Bytes())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if !rows.Next() {
		str := fmt.Sprintf("the given hash (%s) do NOT match a Layer instance", hash.String())
		return nil, errors.New(str)
	}

	var row uint
	var height uint
	err = rows.Scan(&row, &height)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return app.viewportBuilder.Create().
		WithRow(row).
		WithHeight(height).
		Now()
}
