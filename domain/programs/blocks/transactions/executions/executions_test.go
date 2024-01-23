package executions

import (
	"reflect"
	"testing"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	token_layers "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

func TestExecutions_Success(t *testing.T) {
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

	list := []Execution{
		NewExecutionForTests(actionsIns),
	}

	ins := NewExecutionsForTests(list)

	retList := ins.List()
	if !reflect.DeepEqual(list, retList) {
		t.Errorf("the returned execution list is invalid")
		return
	}
}

func TestExecutions_withEmptyList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithList([]Execution{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestExecutions_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}
