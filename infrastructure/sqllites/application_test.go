package sqllites

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"bytes"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/credentials"
	"steve.care/network/domain/programs/blocks/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens"
	"steve.care/network/infrastructure/edwards25519"
	"steve.care/network/infrastructure/jsons"

	token_layers "steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

type resourceExec struct {
	name     string
	resource resources.Resource
}

func TestApplication_Account_InsertThenRetrieve_Success(t *testing.T) {
	dbDir := "./test_files"
	appIns := NewApplication(
		edwards25519.NewEncryptor(),
		jsons.NewAccountAdapter(),
		4096,
		dbDir,
	)

	dbName := "testdb"
	defer func() {
		path := filepath.Join(dbDir, dbName)
		os.Remove(path)
	}()

	// close:
	defer appIns.Close()

	// init with our schema:
	schema := getSchema()
	initAppIns, err := appIns.Init(dbName, schema)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	accAppIns := initAppIns.Accounts()
	username := "rogerCyr"
	password := []byte("this is my password")
	credentials, err := credentials.NewBuilder().Create().WithUsername(username).WithPassword(password).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// insert account:
	err = accAppIns.Insert(credentials)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = appIns.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retAccount, err := accAppIns.Retrieve(credentials)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(username, retAccount.Username()) {
		t.Errorf("the returned account is invalid")
		return
	}

	// delete:
	/*err = accAppIns.Delete(credentials)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = appIns.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	exists, err := accAppIns.Exists(username)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if exists {
		t.Errorf("the account (username: %s) was NOT expected to exists", username)
		return
	}*/
}

func TestApplication_Resources_InsertThenRetrieve_Success(t *testing.T) {
	dbDir := "./test_files"
	appIns := NewApplication(
		edwards25519.NewEncryptor(),
		jsons.NewAccountAdapter(),
		4096,
		dbDir,
	)

	dbName := "testdb"
	defer func() {
		path := filepath.Join(dbDir, dbName)
		os.Remove(path)
	}()

	// close:
	defer appIns.Close()

	// init with our schema:
	schema := getSchema()
	_, err := appIns.Init(dbName, schema)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	firstToken := tokens.NewTokenWithLayerForTests(
		token_layers.NewLayerWithLayerForTests(
			layers.NewLayerForTests(
				layers.NewInstructionsForTests([]layers.Instruction{
					layers.NewInstructionWithStopForTests(),
				}),
				layers.NewOutputForTests(
					"myVariable",
					layers.NewKindWithPromptForTests(),
				),
			),
		),
	)

	firstMsg := firstToken.Hash().Bytes()
	firstSignature, err := signers.NewFactory().Create().Sign(firstMsg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	secondToken := tokens.NewTokenWithLayerForTests(
		token_layers.NewLayerWithLayerForTests(
			layers.NewLayerForTests(
				layers.NewInstructionsForTests([]layers.Instruction{
					layers.NewInstructionWithStopForTests(),
				}),
				layers.NewOutputForTests(
					"myVariable",
					layers.NewKindWithPromptForTests(),
				),
			),
		),
	)

	secondMsg := secondToken.Hash().Bytes()
	secondSignature, err := signers.NewFactory().Create().Sign(secondMsg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// make the resources execution list:
	execList := []resourceExec{
		{
			name: "layerBytesReferenceWithVariable",
			resource: resources.NewResourceForTests(
				firstToken,
				firstSignature,
			),
		},
		{
			name: "layerBytesReferenceWithBytes",
			resource: resources.NewResourceForTests(
				secondToken,
				secondSignature,
			),
		},
	}

	// execute the execution
	for _, oneExec := range execList {
		// begin the transaction:
		coreAppIns, err := appIns.Begin(dbName)
		if err != nil {
			t.Errorf("(execution name: %s) - the error was expected to be nil, error returned: %s", oneExec.name, err.Error())
			return
		}

		// fetch the resource application:
		resAppIns := coreAppIns.Resources()

		// insert resource:
		err = resAppIns.Insert(oneExec.resource)
		if err != nil {
			t.Errorf("(execution name: %s) - the error was expected to be nil, error returned: %s", oneExec.name, err.Error())
			return
		}

		// commit:
		err = appIns.Commit()
		if err != nil {
			t.Errorf("(execution name: %s) - the error was expected to be nil, error returned: %s", oneExec.name, err.Error())
			return
		}

		retResource, err := resAppIns.RetrieveByHash(oneExec.resource.Hash())
		if err != nil {
			t.Errorf("(execution name: %s) - the error was expected to be nil, error returned: %s", oneExec.name, err.Error())
			return
		}

		if !bytes.Equal(retResource.Hash().Bytes(), oneExec.resource.Hash().Bytes()) {
			t.Errorf("(execution name: %s) - the returned resource is invalid", oneExec.name)
			return
		}

		// delete:
		/*secAppIns, err := appIns.BeginInMemory()
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		secAccountAppIns := secAppIns.Accounts()
		err = secAccountAppIns.Delete(credentials)
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}*/

		// commit:
		/*err = appIns.Commit()
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		exists, err := secAccountAppIns.Exists(username)
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		if exists {
			t.Errorf("the account (username: %s) was NOT expected to exists", username)
			return
		}*/
	}

}
