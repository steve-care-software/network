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

func TestTransaction_Success(t *testing.T) {
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

	ins := NewTransactionForTests(executions, execSignature)

	retExecutions := ins.Executions()
	if !reflect.DeepEqual(executions, retExecutions) {
		t.Errorf("the returned executions is invalid")
		return
	}

	retSignature := ins.Signature()
	if !reflect.DeepEqual(execSignature, retSignature) {
		t.Errorf("the returned signature is invalid")
		return
	}
}

func TestTransaction_withoutExecutions_returnsError(t *testing.T) {
	executionsMsg := []byte("this is some msg")
	execSignature, err := signers.NewFactory().Create().Sign(executionsMsg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	_, err = NewTransactionBuilder().Create().WithSignature(execSignature).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestTransaction_withoutSignature_returnsError(t *testing.T) {
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

	_, err = NewTransactionBuilder().Create().WithExecutions(executions).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}
