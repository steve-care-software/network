package transactions

import (
	"reflect"
	"testing"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/programs/blocks/transactions/executions"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	token_layers "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

func TestTransactions_Success(t *testing.T) {
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

	list := []Transaction{
		NewTransactionForTests(executions, execSignature),
	}

	ins := NewTransactionsForTests(list)

	retList := ins.List()
	if !reflect.DeepEqual(list, retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestTransactions_withEmptyList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithList([]Transaction{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestTransactions_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithList([]Transaction{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}
