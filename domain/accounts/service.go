package accounts

import (
	_ "github.com/mattn/go-sqlite3"
	account_encryptors "steve.care/network/domain/accounts/encryptors"
	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/databases/criterias/conditions"
	"steve.care/network/domain/databases/criterias/entities/resources"
	"steve.care/network/domain/databases/transactions"
	"steve.care/network/domain/encryptors"
)

type service struct {
	resourceBuilder          resources.Builder
	conditionBuilder         conditions.Builder
	conditionPointerBuilder  conditions.PointerBuilder
	conditionOperatorBuilder conditions.OperatorBuilder
	conditionElementBuilder  conditions.ElementBuilder
	conditionResourceBuilder conditions.ResourceBuilder
	encryptor                encryptors.Encryptor
	builder                  Builder
	repository               Repository
	adapter                  Adapter
	encryptorBuilder         account_encryptors.Builder
	signerFactory            signers.Factory
	trx                      transactions.Transaction
	bitrate                  int
}

func createService(
	resourceBuilder resources.Builder,
	conditionBuilder conditions.Builder,
	conditionPointerBuilder conditions.PointerBuilder,
	conditionOperatorBuilder conditions.OperatorBuilder,
	conditionElementBuilder conditions.ElementBuilder,
	conditionResourceBuilder conditions.ResourceBuilder,
	encryptor encryptors.Encryptor,
	builder Builder,
	repository Repository,
	adapter Adapter,
	encryptorBuilder account_encryptors.Builder,
	signerFactory signers.Factory,
	trx transactions.Transaction,
	bitrate int,
) Service {
	out := service{
		resourceBuilder:          resourceBuilder,
		conditionBuilder:         conditionBuilder,
		conditionPointerBuilder:  conditionPointerBuilder,
		conditionOperatorBuilder: conditionOperatorBuilder,
		conditionElementBuilder:  conditionElementBuilder,
		conditionResourceBuilder: conditionResourceBuilder,
		encryptor:                encryptor,
		trx:                      trx,
		builder:                  builder,
		repository:               repository,
		adapter:                  adapter,
		encryptorBuilder:         encryptorBuilder,
		signerFactory:            signerFactory,
		bitrate:                  bitrate,
	}

	return &out
}

// Insert inserts an account
func (app *service) Insert(account Account, password []byte) error {
	bytes, err := app.adapter.ToBytes(account)
	if err != nil {
		return err
	}

	cipher, err := app.encryptor.Encrypt(bytes, password)
	if err != nil {
		return err
	}

	username := account.Username()
	if err != nil {
		return err
	}

	return app.trx.Insert("accounts", map[string]any{
		"username": username,
		"cipher":   cipher,
	})
}

// Update updates an account
func (app *service) Update(credentials credentials.Credentials, criteria UpdateCriteria) error {
	originAccount, err := app.repository.Retrieve(credentials)
	if err != nil {
		return err
	}

	originUsername := originAccount.Username()
	originSigner := originAccount.Signer()
	originEncryptor := originAccount.Encryptor()

	builder := app.builder.Create().
		WithUsername(originUsername).
		WithEncryptor(originEncryptor).
		WithSigner(originSigner)

	if criteria.ChangeEncryptor() {
		encryptor, err := app.encryptorBuilder.Create().
			WithBitRate(app.bitrate).
			Now()

		if err != nil {
			return err
		}

		builder.WithEncryptor(encryptor)
	}

	if criteria.ChangeSigner() {
		signer := app.signerFactory.Create()
		builder.WithSigner(signer)
	}

	if criteria.HasUsername() {
		username := criteria.Username()
		builder.WithUsername(username)
	}

	updatedAccount, err := builder.Now()
	if err != nil {
		return err
	}

	bytes, err := app.adapter.ToBytes(updatedAccount)
	if err != nil {
		return err
	}

	updatedPassword := credentials.Password()
	if criteria.HasPassword() {
		updatedPassword = criteria.Password()
	}

	cipher, err := app.encryptor.Encrypt(bytes, updatedPassword)
	if err != nil {
		return err
	}

	return app.trx.Execute("UPDATE accounts set username = ?, cipher = ? where username = ?", updatedAccount.Username(), cipher, originUsername)
}

// Delete deletes an account
func (app *service) Delete(credentials credentials.Credentials) error {
	username := credentials.Username()
	condition, err := app.conditionUsernameEqualsUsername(username)
	if err != nil {
		return err
	}

	resource, err := app.resourceBuilder.Create().
		WithContainer("accounts").
		WithCondition(condition).
		Now()

	if err != nil {
		return err
	}

	return app.trx.Delete(resource)
}

func (app *service) conditionUsernameEqualsUsername(username string) (conditions.Condition, error) {
	usernameField, err := app.conditionPointerBuilder.Create().
		WithContainer("accounts").
		WithField("username").
		Now()

	if err != nil {
		return nil, err
	}

	equalOperator, err := app.conditionOperatorBuilder.Create().
		IsEqual().
		Now()

	if err != nil {
		return nil, err
	}

	usernameResource, err := app.conditionResourceBuilder.Create().
		WithValue(username).
		Now()

	if err != nil {
		return nil, err
	}

	usernameElement, err := app.conditionElementBuilder.Create().
		WithResource(usernameResource).
		Now()

	if err != nil {
		return nil, err
	}

	return app.conditionBuilder.Create().
		WithPointer(usernameField).
		WithOperator(equalOperator).
		WithElement(usernameElement).
		Now()
}
