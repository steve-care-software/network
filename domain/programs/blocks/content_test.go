package blocks

import (
	"reflect"
	"testing"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions"
	"steve.care/network/domain/programs/blocks/transactions/executions"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	token_layers "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

func TestContent_Success(t *testing.T) {
	token := tokens.NewTokenWithLayerForTests(
		token_layers.NewLayerWithLayerForTests(
			layers.NewLayerForTests(
				layers.NewInstructionsForTests([]layers.Instruction{
					layers.NewInstructionWithStopForTests(),
				}),
				layers.NewOutputForTests(
					"myVariable",
					layers.NewKindWithPromptForTests(),
				),
				"myInput",
			),
		),
	)

	msg := token.Hash().Bytes()
	signature, err := signers.NewFactory().Create().Sign(msg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	executions := executions.NewExecutionsForTests([]executions.Execution{
		executions.NewExecutionForTests(
			actions.NewActionsForTests([]actions.Action{
				actions.NewActionWithCreateForTests(
					resources.NewResourceForTests(token, signature),
				),
			}),
		),
	})

	executionsMsg := executions.Hash().Bytes()
	execSignature, err := signers.NewFactory().Create().Sign(executionsMsg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	transactions := transactions.NewTransactionsForTests([]transactions.Transaction{
		transactions.NewTransactionForTests(executions, execSignature),
	})

	ins := NewContentForTests(transactions)

	if ins.HasParent() {
		t.Errorf("the content was expected to NOT contain a parent")
		return
	}

	retTrx := ins.Transactions()
	if !reflect.DeepEqual(transactions, retTrx) {
		t.Errorf("the returned transactions is invalid")
		return
	}
}

func TestContent_withParent_Success(t *testing.T) {
	token := tokens.NewTokenWithLayerForTests(
		token_layers.NewLayerWithLayerForTests(
			layers.NewLayerForTests(
				layers.NewInstructionsForTests([]layers.Instruction{
					layers.NewInstructionWithStopForTests(),
				}),
				layers.NewOutputForTests(
					"myVariable",
					layers.NewKindWithPromptForTests(),
				),
				"myInput",
			),
		),
	)

	msg := token.Hash().Bytes()
	signature, err := signers.NewFactory().Create().Sign(msg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	executions := executions.NewExecutionsForTests([]executions.Execution{
		executions.NewExecutionForTests(
			actions.NewActionsForTests([]actions.Action{
				actions.NewActionWithCreateForTests(
					resources.NewResourceForTests(token, signature),
				),
			}),
		),
	})

	executionsMsg := executions.Hash().Bytes()
	execSignature, err := signers.NewFactory().Create().Sign(executionsMsg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	transactions := transactions.NewTransactionsForTests([]transactions.Transaction{
		transactions.NewTransactionForTests(executions, execSignature),
	})

	pParentHash, err := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := NewContentWithParentForTests(transactions, *pParentHash)

	if !ins.HasParent() {
		t.Errorf("the content was expected to contain a parent")
		return
	}

	retParent := ins.Parent()
	if !reflect.DeepEqual(*pParentHash, retParent) {
		t.Errorf("the returned parent hash is invalid")
		return
	}

	retTrx := ins.Transactions()
	if !reflect.DeepEqual(transactions, retTrx) {
		t.Errorf("the returned transactions is invalid")
		return
	}
}

func TestContent_withoutTransactions_returnsError(t *testing.T) {
	_, err := NewContentBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}
