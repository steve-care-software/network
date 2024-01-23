package executions

import (
	"reflect"
	"testing"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	token_layers "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

func TestExecution_Success(t *testing.T) {
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

	actionsIns := actions.NewActionsForTests([]actions.Action{
		actions.NewActionWithCreateForTests(
			resources.NewResourceForTests(token, signature),
		),
	})
	ins := NewExecutionForTests(actionsIns)

	if ins.HasReceipt() {
		t.Errorf("the execution was expected to NOT contain a receipt")
		return
	}

	retActions := ins.Actions()
	if !reflect.DeepEqual(actionsIns, retActions) {
		t.Errorf("the returned actions resource is invalid")
		return
	}
}

func TestExecution_withReceipt_Success(t *testing.T) {
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
		panic(err)
	}

	actionsIns := actions.NewActionsForTests([]actions.Action{
		actions.NewActionWithCreateForTests(
			resources.NewResourceForTests(token, signature),
		),
	})

	pReceiptHash, err := hash.NewAdapter().FromBytes([]byte("this is some data"))
	if err != nil {
		panic(err)
	}

	ins := NewExecutionWithReceiptForTests(actionsIns, *pReceiptHash)

	if !ins.HasReceipt() {
		t.Errorf("the execution was expected to contain a receipt")
		return
	}

	retReceipt := ins.Receipt()
	if !reflect.DeepEqual(*pReceiptHash, retReceipt) {
		t.Errorf("the returned receipt hash is invalid")
		return
	}

	retActions := ins.Actions()
	if !reflect.DeepEqual(actionsIns, retActions) {
		t.Errorf("the returned actions resource is invalid")
		return
	}
}

func TestExecution_withoutActions_returnsError(t *testing.T) {
	_, err := NewExecutionBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}
