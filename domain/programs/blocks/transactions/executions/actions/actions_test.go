package actions

import (
	"reflect"
	"testing"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens"
	token_layers "steve.care/network/domain/programs/blocks/transactions/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

func TestAction_withList_Success(t *testing.T) {
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

	list := []Action{
		NewActionWithCreateForTests(
			resources.NewResourceForTests(token, signature),
		),
	}

	ins := NewActionsForTests(list)

	retList := ins.List()
	if !reflect.DeepEqual(list, retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestAction_withEmptyList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().WithList([]Action{}).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestAction_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
