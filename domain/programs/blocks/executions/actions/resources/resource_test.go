package resources

import (
	"reflect"
	"testing"

	"steve.care/network/domain/accounts/signers"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/blocks/executions/actions/resources/tokens"
	token_layers "steve.care/network/domain/programs/blocks/executions/actions/resources/tokens/layers"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

func TestResource_Success(t *testing.T) {
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
			),
		),
	)

	msg := token.Hash().Bytes()
	signature, err := signers.NewFactory().Create().Sign(msg)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := NewResourceForTests(token, signature)

	retToken := ins.Token()
	if !reflect.DeepEqual(token, retToken) {
		t.Errorf("the returned token is invalid")
		return
	}

	retSignature := ins.Signature()
	if !reflect.DeepEqual(signature, retSignature) {
		t.Errorf("the returned signature is invalid")
		return
	}
}

func TestResource_withoutToken_returnsError(t *testing.T) {
	pMsg, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	signature, err := signers.NewFactory().Create().Sign(*pMsg)
	if err != nil {
		panic(err)
	}

	_, err = NewBuilder().Create().WithSignature(signature).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestResource_withoutSignature_returnsError(t *testing.T) {
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
			),
		),
	)

	_, err := NewBuilder().Create().WithToken(token).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
