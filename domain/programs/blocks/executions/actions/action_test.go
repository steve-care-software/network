package actions

import (
	"reflect"
	"testing"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/executions/actions/resources"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens"
	token_layers "steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

func TestAction_withCreate_Success(t *testing.T) {
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

	create := resources.NewResourceForTests(token, signature)
	ins := NewActionWithCreateForTests(create)

	if !ins.IsCreate() {
		t.Errorf("the action was expected to contain a create")
		return
	}

	if ins.IsDelete() {
		t.Errorf("the action was expected to NOT contain a delete")
		return
	}

	retCreate := ins.Create()
	if !reflect.DeepEqual(create, retCreate) {
		t.Errorf("the returned create resource is invalid")
		return
	}
}

func TestAction_withDelete_Success(t *testing.T) {
	pDel, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	ins := NewActionWithDeleteForTests(*pDel)

	if ins.IsCreate() {
		t.Errorf("the action was expected to NOT contain a create")
		return
	}

	if !ins.IsDelete() {
		t.Errorf("the action was expected to contain a delete")
		return
	}

	retDel := ins.Delete()
	if !reflect.DeepEqual(*pDel, retDel) {
		t.Errorf("the returned delete hash is invalid")
		return
	}
}

func TestAction_withoutParam_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
